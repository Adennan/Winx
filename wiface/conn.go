package wiface

import "net"

// 定义链接模块的抽象层
type Connection interface {
	// start conn
	Start()
	//
	Stop()
	// 获取当前链接的绑定套接字
	GetTcpConnection() *net.TCPConn
	// 获取当前链接模块的链接ID
	GetConnID() uint32
	// 获取远程客户端的 TCP状态 IP
	RemoteAddr() net.Addr
	// 发送数据
	Send(data []byte) error
}

type HandlerFunc func(*net.TCPConn, []byte, int) error
