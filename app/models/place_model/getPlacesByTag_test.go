package place_model

import (
	"testing"
)

func TestGetPlacesByTags(t *testing.T) {
	// Поиск по одному полю.
	searchString := "фитнес"
	places, errors := GetPlacesByTags(searchString)
	if errors != nil {
		t.Log(errors)
		t.Log("FAIL: cant select places by on tag")
		t.Fail()
	} else {
		//t.Log(places)
		t.Log("select places by tag was successful\n", "found ", len(places), " places")
	}

	searchString = ""
	places, errors = GetPlacesByTags(searchString)
	if errors != nil {
		t.Log(errors)
	} else {
		t.Fail()
		//t.Log(places)
		t.Log("FAIL: select places by tag was successful\n", "found ", len(places), " places")
	}
}
