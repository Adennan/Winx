package wnet

import (
	"Winx/wiface"
	"fmt"
	"net"
)

type Connection struct {
	Conn   *net.TCPConn
	ConnID uint32
	Closed bool
	// 管理连接状态
	ExitChan chan bool

	// 该链接处理的方法router
	Router wiface.IRouter
}

// router的本质仍然是处理回调
func NewConnection(conn *net.TCPConn, id uint32, router wiface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   id,
		Closed:   false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
}

func (c *Connection) Worker() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Printf("ConnID = %d Reader is exit. Remote addr is %s\n", c.ConnID, c.RemoteAddr().String())
	// ? 读取完了就该关闭吗
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}

		//if err := c.Handler(c.Conn, buf, cnt); err != nil {
		//	fmt.Printf("ConnID %d handler is error %v", c.ConnID, err)
		//	break
		//}

		req := &Request{
			Conn: c,
			Data: buf,
		}

		// 从路由中找到注册绑定的Conn对应的router调用
		c.Router.PreHandler(req)
		c.Router.Handler(req)
		c.Router.PostHandler(req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start ConnID = ", c.ConnID)
	// 启动读数据的业务
	go c.Worker()
	// TODO 启动写数据的业务
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop ConnID = ", c.ConnID)

	if c.Closed {
		return
	}
	c.Closed = true

	// 回收资源
	c.Conn.Close()

	close(c.ExitChan)
}

// 获取当前链接的绑定套接字
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的 TCP状态 IP
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}
