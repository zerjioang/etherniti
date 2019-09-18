package httpawn

import "net"

const (
	notfound = `not found`
)

var (
	notfoundRaw = []byte(notfound)
)

type Router struct {
	store map[string]HttpOperation
}

func NewRouter() Router {
	return Router{
		store: make(map[string]HttpOperation),
	}
}

func (r Router) Execute(clientAddr net.Addr, m httpMethod, path string) []byte {
	//look for requested operation in the map
	op, ok := r.store[path]
	if ok && op != nil {
		ctx := new(Context)
		ctx.clientAddr = clientAddr
		// execute requested function with given context
		op(ctx)
		// and get serialized body result
		return ctx.Bytes()
	}
	return notfoundRaw
}
