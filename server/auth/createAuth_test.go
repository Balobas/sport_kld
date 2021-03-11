package auth

import (
	"fmt"
	"sport_kld/server/auth/token"
	"testing"
)

func TestCreateAuth(t *testing.T) {
	td, _ := token.CreateToken("fa27e2b4-8189-11eb-956d-0c9d9244")
	err := CreateAuth("fa27e2b4-8189-11eb-956d-0c9d9244", td)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	fmt.Println(td.AccessToken)
	fmt.Println("refresh:")
	fmt.Println(td.RefreshToken)
	t.Log("passed")
}
