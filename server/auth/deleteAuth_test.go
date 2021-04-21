package auth

import "testing"

func TestDeleteAuth(t *testing.T) {
	err := DeleteAuth("c3ec0d7e-5fbe-4016-971c-ab47a455b1a1")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	err = DeleteAuth("sdfsfd")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}
