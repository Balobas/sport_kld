package place_model

import (
	"../../models"
	"../place_model"
	"testing"
)

func TestGetPlacesByOrganizationUID(t *testing.T) {
	testUid := models.UID("fdc1e5df5f0111eb913f0c9d92446328")

	places, errors := place_model.GetPlacesByOrganizationUID(testUid)
	if len(errors) != 0 {
		t.Log("errors: ")
		t.Log(errors)
	}
	t.Log(places)
	t.Log("found ", len(places), " places")
}
