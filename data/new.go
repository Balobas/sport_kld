package data

import "../core"

func (org Organization) New() core.Object {
	return Organization{
		UID:            "",
		Name:           "",
		Description:    "",
		PlacesUIDs:     nil,
		CategoriesUIDs: nil,
		TagsUIDs:       nil,
	}
}

func (place Place) New() core.Object {
	return Place{
		UID:                "",
		Name:               "",
		BuildingName:       "",
		Description:        "",
		Adress:             "",
		City:               "",
		OpeningHours:       "",
		PostIndex:          "",
		WebSite:            "",
		Phones:             "",
		CategoriesUIDs:     nil,
		TagsUIDs:           nil,
		HolderOrganization: Organization{},
		BasedOrganizations: nil,
		FreeVisit:          false,
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
