package auth

import "testing"

func Test_getAuth(t *testing.T) {
	auth, err := getAuth("77fb2198-b2ca-4c3b-bf9f-67e5515d0e34")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(auth)

	_, err = getAuth("adsad")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}

func Test_getAccess(t *testing.T) {
	a, err := getAccess("75bc1537-cc05-4217-a715-bd7930b3f37e")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(a)

	_, err = getAccess("adsad")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}

func Test_getRefresh(t *testing.T) {
	a, err := getRefresh("5579d170-289f-46b4-a8be-65b26e37a20c")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(a)

	_, err = getRefresh("adsad")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}

func Test_getAuthByUserUid(t *testing.T) {
	a, err := getAuthByUserUid("fa27e2b4-8189-11eb-956d-0c9d9244")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(a)

	_, err = getAuthByUserUid("adsad")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}

func TestAuthExist(t *testing.T) {
	if err := ExistAuth("fa27e2b4-8189-11eb-956d-0c9d9244"); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log("passed")
}

func TestGetAuthUidByAccessUid(t *testing.T) {
	authUid, err := GetAuthUidByAccessUid("a6d95e1d-6a80-46e1-945e-75f33681f95e")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(authUid)

	_, err = GetAuthUidByAccessUid("75bc1")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}

func TestGetAuthUidByRefreshUid(t *testing.T) {
	authUid, err := GetAuthUidByRefreshUid("b45e8448-a57e-48a0-810a-32e368988ff4")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log(authUid)

	_, err = GetAuthUidByRefreshUid("75bc1")
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}
