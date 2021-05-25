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
		CreatorUID:    "",
		IsPrivate:     false,
		EventPassword: "",
		IsOver:        false,
	}

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	t := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdWlkIjoiNjM2YmZhOWMtOWUyYy00Y2I5LWJlNzUtY2FjZTBiNmU2N2IwIiwiZXhwIjoxNjIwMzM1MTU4LCJ1c2VyX3VpZCI6IjhlY2Q4OWU0LWFlYWEtMTFlYi04OGY3LTdlYjI3MjBhZGM3MSJ9.8f8UgD8pGLwxbc6qX7jW2ommEtxi8n84cqmdz4VRFgM"

	client := &http.Client{}
	URL := config.BaseUrl + "/event"
	req, err := http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", t))
	resp, err := client.Do(req)
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
}
