package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sport_kld/app/models/event_model"
	"sport_kld/test_app_api/config"
)

func TestCreateEvent() {
	params := event_model.Event{
		UID:           "",
		Name:          "event123",
		Description:   "123123123",
		Dates:         "12-01-12 14-41-12",
		Time:          "12:30",
		VisitorsNum:   0,
		VisitorsLimit: 5,
		PlaceUID:      "fe07e92c5f0111ebb0050c9d92446328",
		CreatorUID:    "fa27e2b4-8189-11eb-956d-0c9d9244",
		IsPrivate:     false,
		EventPassword: "",
		IsOver:        false,
	}

	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(config.BaseUrl + "/event", "application/json", r)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("DONE: CreateEvent test")
}
