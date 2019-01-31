package profile

import "testing"

func TestCreateConnectionProfileToken(t *testing.T) {
	p := ConnectionProfile{}
	token, err := CreateConnectionProfileToken(p)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
