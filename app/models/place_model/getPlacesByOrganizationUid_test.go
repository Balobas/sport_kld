package place_model

import (
	"sport_kld/app/models"
	"testing"
)

func TestGetPlacesByOrganizationUID(t *testing.T) {
	testUid := models.UID("fdc1e5df5f0111eb913f0c9d92446328")

	places, errors := GetPlacesByOrganizationUID(testUid)
	if len(errors) != 0 {
		t.Log("errors: ")
		t.Log(errors)
	}
	t.Log(places)
	t.Log("found ", len(places), " places")
}
