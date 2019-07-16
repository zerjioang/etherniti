package project

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectModel(t *testing.T) {
	t.Run("create-project-model", func(t *testing.T) {
		p := NewProject("random-name", nil)
		assert.NotNil(t, p)
		assert.Equal(t, p.Name, "random-name")
		assert.NotEqual(t, p.ProjectId, "")
		assert.NotEqual(t, p.ProjectSecret, "")
		t.Log(p)
	})
	t.Run("deserialize", func(t *testing.T) {
		p := NewProject("random-name", nil)
		v := p.Value(testSerializer)
		p2 := NewEmptyProject()
		err := json.Unmarshal(v, &p2)
		assert.Nil(t, err)
		assert.Equal(t, p2.Name, "random-name")
	})
}
