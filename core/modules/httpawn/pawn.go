package httpawn

// callback operations that specify the code
// to be executed on each http call
type HttpOperation func(c *Context)

// http pawned server
type PawnServer struct {
	r Router
}

func New() PawnServer {
	return PawnServer{
		r: NewRouter(),
	}
}

func (server *PawnServer) GET(path string, operation HttpOperation) {
	server.r.store[path] = operation
}

func (server *PawnServer) Start(addr string) error {
	return Serve(addr, &server.r)
}
