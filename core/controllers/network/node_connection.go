package network

type NodeConnection struct {
	// peer connection protocol: http, https, unix
	protocol string
	//main connection peer address/ip
	ip string
	// rpc rpcPort
	rpcPort string
	// graphql rpcPort
	graphqlPort string
	// websocket rpcPort
	wsPort string
	// additional URI content after port definition
	path string
	// connection name: mainet, ropsten, rinkeby, custom, etc
	name string
	//raw uri string
	raw string
}

func (c NodeConnection) Protocol() string {
	return c.protocol
}

func (c NodeConnection) Ip() string {
	return c.ip
}

func (c NodeConnection) Port() string {
	return c.rpcPort
}

func (c NodeConnection) Name() string {
	return c.name
}

func (c NodeConnection) GetRPCEndpoint() string {
	// if there is a raw URI defined, just return the URI
	if c.raw != "" {
		return c.raw
	} else {
		if c.rpcPort == "" {
			//no rpc port specified. assuming 80 and 443
			return c.protocol + c.ip + c.path
		} else {
			// rpc port specified
			return c.protocol + c.ip + ":" + c.rpcPort + c.path
		}
	}
}

func (c NodeConnection) GetWebsocketEndpoint() string {
	// if there is a raw URI defined, just return the URI
	if c.raw != "" {
		return c.raw
	} else {
		if c.wsPort == "" {
			//no websocket port specified. assuming 80 and 443
			return c.protocol + c.ip + c.path
		} else {
			// websocket port specified
			return c.protocol + c.ip + ":" + c.wsPort + c.path
		}
	}
}

func (c NodeConnection) GetGraphQLEndpoint() string {
	// if there is a raw URI defined, just return the URI
	if c.raw != "" {
		return c.raw
	} else {
		if c.graphqlPort == "" {
			//no graphql port specified. assuming 80 and 443
			return c.protocol + c.ip + c.path
		} else {
			// graphql port specified
			return c.protocol + c.ip + ":" + c.graphqlPort + c.path
		}
	}
}

func NewNodeConnection(protocol, ip, port, graphqlPort, wsPort, path, name string) *NodeConnection {
	nc := new(NodeConnection)
	nc.protocol = protocol
	nc.ip = ip
	nc.rpcPort = port
	nc.graphqlPort = graphqlPort
	nc.wsPort = wsPort
	nc.path = path
	nc.name = name
	return nc
}

func NewUndefinedConnection() *NodeConnection {
	return new(NodeConnection)
}

func NodeConnectionFromString(nodeURI string) *NodeConnection {
	nc := new(NodeConnection)
	nc.raw = nodeURI
	return nc
}
