package organization_model

import (
	"sport_kld/app/models"
	"testing"
)

func TestGetOrganizationByPlaceUID(t *testing.T) {
	testUid := models.UID("fdc1e5e85f0111eb828e0c9d92446328")

	orgs, errors := GetOrganizationByPlaceUID(testUid)
	if len(errors) != 0 {
		t.Log("errors: ")
		t.Log(errors)
	}
	t.Log(orgs)
	t.Log("found ", len(orgs), " organization(s)")
}
