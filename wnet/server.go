package wnet

import (
	"Winx/wiface"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    wiface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[]")

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		lis, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}

		fmt.Println("Start Winx server success")
		var cid uint32 = 0

		for {
			conn, err := lis.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//// 已与客户端建立连接
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("recv buf err: ", err)
			//			continue
			//		}
			//
			//		// 回显功能
			//		if _, err := conn.Write(buf[:cnt]); err != nil {
			//			fmt.Println("write back buf err: ", err)
			//			continue
			//		}
			//	}
			//}()

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(r wiface.IRouter) {
	s.Router = r
	fmt.Println("Add Router Success!")
}

func NewServer(name string) wiface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "ipv4",
		IP:        "0.0.0.0",
		Port:      9999,
		Router:    nil,
	}
}

// 定义默认的客户端链接所绑定的 handler api
func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显
	fmt.Println("[Conn Handler] CallbackToClient")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallbackToClient error")
	}

	return nil
}
