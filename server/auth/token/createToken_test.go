package token

import "testing"

func TestCreateToken(t *testing.T) {
	td, err := CreateToken("fa27e2b4-8189-11eb-956d-0c9d9244")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(td)
	t.Log("passed")
}
