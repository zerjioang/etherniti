package project

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/util/ip"
)

func TestProjectModel(t *testing.T) {
	t.Run("create-project-model", func(t *testing.T) {
		testIp := ip.Ip2intLow("127.0.0.1")
		p := NewProject("random-name", "jhon.doe@domain.tld", testIp)
		assert.NotNil(t, p)
		assert.Equal(t, p.Name, "random-name")
		assert.Equal(t, p.Owner, "jhon.doe@domain.tld")
		assert.NotEqual(t, p.ProjectId, "")
		assert.NotEqual(t, p.ProjectSecret, "")
		assert.True(t, p.CreationDate > 0)
		t.Log(p)
	})
	t.Run("deserialize", func(t *testing.T) {
		testIp := ip.Ip2intLow("127.0.0.1")
		p := NewProject("random-name", "jhon.doe@domain.tld", testIp)
		_, v := p.Storage()
		p2 := NewEmptyProject()
		err := json.Unmarshal(v, &p2)
		assert.Nil(t, err)
		assert.Equal(t, p2.Ip, uint32(2130706433))
		assert.Equal(t, p2.Name, "random-name")
		assert.Equal(t, p2.Owner, "jhon.doe@domain.tld")
	})
}
