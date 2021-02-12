package organization_model

import "testing"

func TestGetOrganizationsByFields(t *testing.T) {
	// Поиск по одному полю.
	searchString := "Альбатрос"
	fieldsNames := []string{"name"}

	orgs, errors := GetOrganizationsByFields(fieldsNames, searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select organization by on field")
		t.Fail()
	} else {
		//t.Log(orgs)
		t.Log("select organizations by one field was successful\n", "found ", len(orgs), " organizations")
	}

	// Поиск по нескольким полям.
	searchString = "Баскетбол"
	fieldsNames = []string{"name", "description"}

	orgs, errors = GetOrganizationsByFields(fieldsNames, searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select organizations by two fields")
		t.Fail()
	} else {
		//t.Log(orgs)
		t.Log("select organizations by one two fields was successful\n", "found ", len(orgs), " organizations")
	}

	// Поиск без указания поля.
	searchString = "Футбол"
	fieldsNames = []string{}

	orgs, errors = GetOrganizationsByFields(fieldsNames, searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select organizations by empty fields array")
		t.Fail()
	} else {
		//t.Log(orgs)
		t.Log("select organizations by empty fields array was successful\n", "found ", len(orgs), " organizations")
	}
}

