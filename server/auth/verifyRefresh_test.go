package auth

import "testing"

func TestVerifyRefresh(t *testing.T) {
	authUid, userUid, err := VerifyRefresh("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NDE3NzQsInJlZnJlc2hfdWlkIjoiMmViYTgwZGMtMzE3ZS00NjcxLWFlNWQtNWQyN2VlMjc5ZjVhIiwidXNlcl91aWQiOiJmYTI3ZTJiNC04MTg5LTExZWItOTU2ZC0wYzlkOTI0NCJ9.pIeZDaxMZw2SVw5ysssIXnxw0nIzUtahH-ndRPM3_wE")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(authUid)
	t.Log(userUid)

	authUid, userUid, err = VerifyRefresh("td.AccessToken")
	if err == nil {
		t.Log("expected error")
		t.FailNow()
	}

	t.Log("passed")
}
