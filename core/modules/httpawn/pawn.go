package httpawn

// callback operations that specify the code
// to be executed on each http call
type HttpOperation func(c *Context)

// http pawned server
type PawnServer struct {
	// server configuration options
	opts PawnOptions
	// server router
	r Router
}

func New(opts ...PawnOptions) PawnServer {
	var s PawnServer
	if opts == nil || len(opts) == 0 {
		s = PawnServer{
			opts: DefaultOptions,
			r:    NewRouter(),
		}
	} else {
		s = PawnServer{
			opts: opts[0],
			r:    NewRouter(),
		}
	}
	// apply server configuration
	// according to specified options
	s.configure()
	return s
}

func (server *PawnServer) GET(path string, operation HttpOperation) {
	server.r.store[path] = operation
}

func (server *PawnServer) Start(addr string) error {
	return Serve(addr, &server.r)
}

// apply to current server instance defined configuration
func (server *PawnServer) configure() {

}
