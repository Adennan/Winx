package wiface

type IServer interface {
	// start server
	Start()
	Stop()
	Serve()

	// 给当前的服务注册一个路由方法
	// 供客户端的链接处理使用
	AddRouter(r IRouter)
}
