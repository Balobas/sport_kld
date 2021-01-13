package main

import (
	"../core"
	"../data"
	"fmt"
)

func main() {

	place := data.Place{
		UID:            "b2690eaf-0b73-4100-8a14-ebac27ae027c",
		Name:           "MESTi",
		Description:    "assdsd",
		HolderOrganization:   data.Organization{},
		Abonements:     nil,
		Location:       data.LocationInfo{},
		CategoriesUIDs: nil,
		TagsUIDs:       nil,
	}

	uid, err := core.ActionUpdate(place)
	fmt.Println(uid, err)
}

