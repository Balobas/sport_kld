package place_model

import (
	"../../models"
	"../place_model"
	"testing"
)

func TestGetByUID(t *testing.T) {
	testUid := models.UID("35cc8cde5f0111eb960b0c9d92446328")

	place, err := place_model.GetByUID(testUid)
	if err != nil {
		t.Log("FAIL: cant get place by uid")
		t.Fail()
	} else {
		t.Log(place)
	}
}

func TestGetPlacesByUIDs(t *testing.T) {
	testUids := []models.UID{"35ccdac45f0111eb86670c9d92446328", "35cc8cde5f0111eb960b0c9d92446328"}

	places, errors := place_model.GetPlacesByUIDs(testUids)
	if len(places) != 2 {
		t.Log("FAIL: cant get all places by input array")
		t.Log(errors)

		t.Fail()
	}
	t.Log(places)
}