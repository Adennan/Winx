package wiface

type IServer interface {
	// start server
	Start()
	Stop()
	Serve()
}
