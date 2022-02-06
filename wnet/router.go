package wnet

import "Winx/wiface"

type BaseRouter struct {
}

func (br *BaseRouter) PreHandler(r wiface.IRequest)  {}
func (br *BaseRouter) Handler(r wiface.IRequest)     {}
func (br *BaseRouter) PostHandler(r wiface.IRequest) {}
