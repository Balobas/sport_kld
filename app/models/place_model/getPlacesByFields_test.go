package place_model

import (
	"testing"
)

func TestGetPlacesByFields(t *testing.T) {
	// Поиск по одному полю.
	searchString := "Альбатрос"
	fieldsNames := []string{"name"}

	places, errors := GetPlacesByFields(fieldsNames, searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select place by on field")
		t.Fail()
	} else {
		//t.Log(places)
		t.Log("select places by one field was successful\n", "found ", len(places), " places")
	}

	// Поиск по нескольким полям.
	searchString = "Баскетбол"
	fieldsNames = []string{"name", "description"}

	places, errors = GetPlacesByFields(fieldsNames, searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select places by two fields")
		t.Fail()
	} else {
		//t.Log(places)
		t.Log("select places by one two fields was successful\n", "found ", len(places), " places")
	}

	// Поиск без указания поля.
	searchString = "Футбол"
	fieldsNames = []string{}

	places, errors = GetPlacesByFields(fieldsNames, searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select places by empty fields array")
		t.Fail()
	} else {
		//t.Log(places)
		t.Log("select places by empty fields array was successful\n", "found ", len(places), " places")
	}
}
