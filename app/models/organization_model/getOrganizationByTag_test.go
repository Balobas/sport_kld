package organization_model

import "testing"

func TestGetOrganizationsByTags(t *testing.T) {
	// Поиск по одному полю.
	searchString := "фитнес"
	orgs, errors := GetOrganizationsByTags(searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select organizations by on tag")
		t.Fail()
	} else {
		//t.Log(orgs)
		t.Log("select organizations by tag was successful\n", "found ", len(orgs), " organizations")
	}

	searchString = ""
	orgs, errors = GetOrganizationsByTags(searchString)
	if errors != nil {
		t.Log(errors)
	} else {
		t.Fail()
		t.Log("FAIL: select organizations by tag was successful\n", "found ", len(orgs), " organizations")
	}
}
