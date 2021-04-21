package token

import "testing"

func TestParseAccessToken(t *testing.T) {
	td, err := CreateToken("fa27e2b4-8189-11eb-956d-0c9d9244")

	a, err := ParseAccessToken(td.AccessToken)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(a)
	t.Log("passed")
}

func TestParseRefreshToken(t *testing.T) {
	td, err := CreateToken("fa27e2b4-8189-11eb-956d-0c9d9244")

	a, err := ParseRefreshToken(td.RefreshToken)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(a)
	t.Log("passed")
}
