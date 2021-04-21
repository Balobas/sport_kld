package auth

import (
	"testing"
)

func TestVerifyAuth(t *testing.T) {
	authUid, userUid, err := VerifyAuth("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdWlkIjoiNTY2YzljN2UtOTgxYy00ODIzLWI0ODQtN2YyMTFjOWQ2ODA1IiwiZXhwIjoxNjE1NDcwNjM5LCJ1c2VyX3VpZCI6ImZhMjdlMmI0LTgxODktMTFlYi05NTZkLTBjOWQ5MjQ0In0.95aGo5FZ-L6pGucnzd4vM_0z0yaD1OkltAhWlqVL8hU")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(authUid)
	t.Log(userUid)

	authUid, userUid, err = VerifyAuth("td.AccessToken")
	if err == nil {
		t.Log("expected error")
		t.FailNow()
	}

	t.Log("passed")
}
