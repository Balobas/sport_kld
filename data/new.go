package data

import "../core"

func (org Organization) New() core.Object {
	return Organization{
		UID:            "",
		Name:           "",
		MainAddress:    "",
		Abonements:     nil,
		Places:         nil,
		CategoriesUIDs: nil,
		TagsUIDs:       nil,
	}
}

func (place Place) New() core.Object {
	return Place{
		UID:                "",
		Name:               "",
		Description:        "",
		HolderOrganization: Organization{},
		BasedOrganizations: nil,
		FreeVisit:          false,
		Abonements:         nil,
		Location:           LocationInfo{},
		CategoriesUIDs:     nil,
		TagsUIDs:           nil,
	}
}

func (abonement Abonement) New() core.Object {
	return Abonement{
		UID:          "",
		Name:         "",
		Description:  "",
		Price:        "",
		ActivateTime: "",
	}
}
