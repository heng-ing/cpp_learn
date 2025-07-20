package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	addr            = ":6000"
	heartbeatPeriod = 60 * 10 * time.Second
	adminPwd        = "secret123"
)

// ---------- 数据结构 ----------

type Message struct {
	From, To, Content string
	Time              time.Time
	Private           bool
}

type Cmd struct {
	Action, Target string
	Source         *Client
}

type Client struct {
	conn    net.Conn
	name    string
	send    chan Message
	alive   chan struct{}
	isAdmin bool
}

type Hub struct {
	clients    map[string]*Client
	mu         sync.RWMutex
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
	commands   chan Cmd
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message, 64),
		commands:   make(chan Cmd, 16),
	}
}

// ---------- Hub 主循环 ----------

func (h *Hub) run(ctx context.Context) {
	for {
		select {
		case c := <-h.register:
			h.mu.Lock()
			h.clients[c.name] = c
			h.mu.Unlock()
			h.broadcast <- Message{From: "System", Content: fmt.Sprintf("%s 加入了聊天室", c.name), Time: time.Now()}
		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c.name]; ok {
				delete(h.clients, c.name)
				close(c.send)
				_ = c.conn.Close()
				h.broadcast <- Message{From: "System", Content: fmt.Sprintf("%s 离开了聊天室", c.name), Time: time.Now()}
			}
			h.mu.Unlock()
		case m := <-h.broadcast:
			h.mu.RLock()
			for _, cli := range h.clients {
				if m.Private && cli.name != m.To && cli.name != m.From {
					continue
				}
				select {
				case cli.send <- m:
				default:
					// 防止写阻塞导致 Hub 卡死
					go func(cl *Client) { h.unregister <- cl }(cli)
				}
			}
			h.mu.RUnlock()
		case cmd := <-h.commands:
			switch cmd.Action {
			case "kick":
				h.mu.RLock()
				target, ok := h.clients[cmd.Target]
				h.mu.RUnlock()
				if ok {
					target.send <- Message{From: "System", Content: "你已被管理员踢出", Time: time.Now()}
					h.unregister <- target
					cmd.Source.send <- Message{From: "System", Content: fmt.Sprintf("已踢出 %s", cmd.Target), Time: time.Now()}
				} else {
					cmd.Source.send <- Message{From: "System", Content: "用户不存在", Time: time.Now()}
				}
			case "list":
				h.mu.RLock()
				names := make([]string, 0, len(h.clients))
				for n := range h.clients {
					names = append(names, n)
				}
				h.mu.RUnlock()
				cmd.Source.send <- Message{From: "System", Content: "在线用户:" + strings.Join(names, ","), Time: time.Now()}
			}
		case <-ctx.Done():
			log.Info("Hub shutting down")
			return
		}
	}
}

// ---------- Client goroutines ----------

func (c *Client) readPump(h *Hub) {
	defer func() { h.unregister <- c }()
	reader := bufio.NewScanner(c.conn)
	for {
		c.conn.SetReadDeadline(time.Now().Add(heartbeatPeriod * 2))
		if !reader.Scan() {
			return
		}
		text := strings.TrimSpace(reader.Text())
		if text == "" {
			continue
		}
		if text == "/ping" {
			c.alive <- struct{}{}
			continue
		}
		if strings.HasPrefix(text, "/") {
			parseCommand(text, c, h)
			continue
		}
		h.broadcast <- Message{From: c.name, Content: text, Time: time.Now()}
	}
}

func parseCommand(s string, c *Client, h *Hub) {
	args := strings.SplitN(s, " ", 3)
	switch args[0] {
	case "/w":
		if len(args) < 3 {
			c.send <- Message{From: "System", Content: "用法: /w <用户> <内容>", Time: time.Now()}
			return
		}
		h.broadcast <- Message{From: c.name, To: args[1], Content: args[2], Private: true, Time: time.Now()}
	case "/kick":
		if !c.isAdmin {
			c.send <- Message{From: "System", Content: "权限不足", Time: time.Now()}
			return
		}
		if len(args) < 2 {
			c.send <- Message{From: "System", Content: "用法: /kick <用户>", Time: time.Now()}
			return
		}
		h.commands <- Cmd{Action: "kick", Target: args[1], Source: c}
	case "/list":
		h.commands <- Cmd{Action: "list", Source: c}
	case "/admin":
		if len(args) == 2 && args[1] == adminPwd {
			c.isAdmin = true
			c.send <- Message{From: "System", Content: "管理员登录成功", Time: time.Now()}
		} else {
			c.send <- Message{From: "System", Content: "密码错误", Time: time.Now()}
		}
	case "/quit":
		h.unregister <- c
	default:
		c.send <- Message{From: "System", Content: "未知命令", Time: time.Now()}
	}
}

func (c *Client) writePump() {
	for m := range c.send {
		line := fmt.Sprintf("[%s %s]%s\n", m.From, m.Time.Format("15:04:05"), m.Content)
		if m.Private {
			line = fmt.Sprintf("[私聊]%s", line)
		}
		c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if _, err := c.conn.Write([]byte(line)); err != nil {
			return
		}
	}
}

// ---------- 主函数 ----------

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	hub := newHub()
	ctx, cancel := context.WithCancel(context.Background())
	go hub.run(ctx)

	// 信号处理
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, os.Kill)
		<-sig
		cancel()
		_ = l.Close()
	}()

	log.Infof("Listening on %s", addr)
	for {
		conn, err := l.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				log.Info("Server stopped")
				return
			default:
				log.Error(err)
				continue
			}
		}
		go handleConn(conn, hub)
	}
}

func handleConn(conn net.Conn, h *Hub) {
	_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	_, _ = conn.Write([]byte("请输入昵称: "))
	reader := bufio.NewScanner(conn)
	if !reader.Scan() {
		_ = conn.Close()
		return
	}
	name := strings.TrimSpace(reader.Text())
	if name == "" {
		_ = conn.Close()
		return
	}
	h.mu.Lock()
	if _, exists := h.clients[name]; exists {
		// _ = conn.Write([]byte("昵称已被占用\n"))
		_, err := conn.Write([]byte("昵称已被占用\n"))
		if err != nil {
			log.Printf("发送昵称占用提示失败: %v", err)
		}
		h.mu.Unlock()
		_ = conn.Close()
		return
	}
	h.mu.Unlock()

	c := &Client{conn: conn, name: name, send: make(chan Message, 16), alive: make(chan struct{}, 1)}
	h.register <- c
	go c.writePump()
	c.readPump(h)
}
