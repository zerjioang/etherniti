package httpawn

import "net"

type Context struct {
	// ip address of the client that makes the request
	clientAddr net.Addr

	// raw http content defined by users
	payload []byte
}

func (c *Context) String(str string) {
	c.payload = StringToBytes(str)
}

func (c *Context) Bytes() []byte {
	return c.payload
}
