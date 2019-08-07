package network

import (
	"errors"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth web3 graphql controller
type GraphqlController struct {
	network *NetworkController
}

// constructor like function
func NewGraphqlController(network *NetworkController) GraphqlController {
	ctl := GraphqlController{}
	ctl.network = network
	return ctl
}

// profile abi data getter
func (ctl *GraphqlController) query(c *echo.Context) error {
	req := c.Body()
	if req == nil || len(req) == 0 {
		// return a binding error
		logger.Error(data.FailedToBind, errors.New("invalid GraphQL query provided"))
		return api.ErrorBytes(c, data.BindErr)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// run requested query against eth node
	result, err := client.GraphQL(ctl.network.GetGraphQLEndpoint(), req)
	if err.None() {
		return api.SendSuccessBlob(c, result)
	} else {
		//grapql query error
		return api.Error(c, err)
	}
}

// implemented method from interface RouterRegistrable
func (ctl GraphqlController) RegisterRouters(router *echo.Group) {
	logger.Debug("adding controller raw GRAPHQL call supports")
	logger.Debug("exposing POST /graphql")
	router.POST("/ql", ctl.query)
}
