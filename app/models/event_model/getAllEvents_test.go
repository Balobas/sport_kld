package event_model

import (
	"fmt"
	"testing"
)

func TestGetAllEvents(t *testing.T) {
	res, err := GetAllEvents()
	fmt.Println(res, err)
}
