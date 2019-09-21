package httpawn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zerjioang/etherniti/core/modules/httpawn"
)

func TestPawnServer(t *testing.T) {
	t.Run("simple-get", func(t *testing.T) {
		server := httpawn.New()
		server.GET("/", func(ctx *httpawn.Context) {
			ctx.String("Hello World!")
		})
		err := server.Start(":8080")
		assert.NoError(t, err)
	})
}
