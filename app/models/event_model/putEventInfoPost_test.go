package event_model

import "testing"

func TestPutEventInfoPost(t *testing.T) {
	uid, err := PutEventInfoPost(EventInfoPost{
		UID:         "",
		EventUID:    "1491b99d-7778-11eb-8979-0c9d9244",
		AuthorUID:   "ttt",
		Text:        "lol",
		CreatedTime: "",
		UpdatedTime: "",
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(uid)
	t.Log("passed")
}
