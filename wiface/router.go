package wiface

type IRouter interface {
	PreHandler(r IRequest)
	Handler(r IRequest)
	PostHandler(r IRequest)
}
