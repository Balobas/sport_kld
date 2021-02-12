package organization_model

import (
	"sport_kld/app/models"
	"testing"
)

func TestGetOrganizationByUID(t *testing.T) {
	testUid := models.UID("fdc25bec5f0111ebba300c9d92446328")

	place, err := GetOrganizationByUID(testUid)
	if err != nil {
		t.Log("FAIL: cant get organization by uid", err)
		t.Fail()
	} else {
		t.Log(place)
	}
}

func TestGetOrganizationsByUIDs(t *testing.T) {
	testUids := []models.UID{"fdc2f8165f0111eba4a90c9d92446328", "fdc456af5f0111ebb45c0c9d92446328"}

	places, errors := GetOrganizationsByUIDs(testUids)
	if len(places) != 2 {
		t.Log("FAIL: cant get all organizations by input array")
		t.Log(errors)

		t.Fail()
	}
	t.Log(places)
}