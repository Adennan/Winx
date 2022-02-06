package wiface

// 将客户端请求的链接信息和请求的数据封装在一起

type IRequest interface {
	GetConn() IConnection
	GetData() []byte
}
