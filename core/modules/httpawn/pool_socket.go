package httpawn

import (
	"net"
	"sync"
)

type socketRequest struct {
	client net.Conn
	raw    []byte
}

type socketRequestPool struct {
	pool sync.Pool
}

var (
	// this is our http server request pool
	sPool socketRequestPool
)

func init() {
	initPool()
}

// Func to init pool
func initPool() {
	sPool = socketRequestPool{}
	sPool.pool = sync.Pool{}
	sPool.pool.New = func() interface{} {
		return new(socketRequest)
	}
}

func (p *socketRequestPool) Load(conn net.Conn, bytes []byte) *socketRequest {
	req := p.pool.Get().(*socketRequest)
	req.client = conn
	req.raw = bytes
	return req
}

func (p *socketRequestPool) Store(req *socketRequest) {
	req.client = nil
	req.raw = nil
	p.pool.Put(req)
}
