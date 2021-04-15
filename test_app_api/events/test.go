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
	event := event_model.Event{
		UID:           "123",
		Name:          "test event",
		Description:   "test event description",
		Dates:         "25.04.2021",
		Time:          "13:00",
		VisitorsNum:   0,
		VisitorsLimit: 25,
		PlaceUID:      "8912829312312312",
		CreatorUID:    "ttt",
		IsPrivate:     false,
		EventPassword: "",
		IsOver:        false,
	}

	b, err := json.Marshal(event)
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
