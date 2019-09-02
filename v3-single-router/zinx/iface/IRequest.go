package iface

type IRequest interface {
	GetIConn() IConnect
	GetData() []byte
	GetDataLen() uint32
}
