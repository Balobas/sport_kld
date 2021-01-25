package data

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type UID string

func (uid UID) String() string {
	return string(uid)
}

func (uid UID) isCorrect() bool {
	return true
}

type Tag struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}

type SportCategory struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}

type Organization struct {
	UID            UID         `json:"uid"`
	Name           string      `json:"name"`
	Description    string 	   `json:"description"`
	PlacesUIDs     []UID       `json:"placesUIDs"`
	TagsUIDs       []UID       `json:"tagsUIDs"`
}

type Place struct {
	UID                UID            `json:"uid"`
	Name               string         `json:"name"`
	BuildingName		string			`json:"buildingName"`
	BuildingType 		string `json:"buildingType"`
	Description        string         `json:"description"`
	Address				string 		`json:"adress"`
	City 				string		`json:"city"`
	OpeningHours		string 		`json:"openingHours"`
	PostIndex			string 		`json:"postIndex"`
	WebSite				string		`json:"webSite"`
	Phones				string		`json:"phones"`
	Email 				string    	`json:"email"`
	Facebook 		string `json:"facebook"`
	Instagram	string `json:"instagram"`
	Twitter string `json:"twitter"`
	VK string `json:"vk"`
	
	TagsUIDs           []UID          `json:"tagsUIDs"`
	//под вопросом//
	HolderOrganization Organization   `json:"organization"`
	BasedOrganizations []Organization `json:"basedOrganizations"`
	FreeVisit          bool           `json:"freeVisit"`
	//
}

//пока под вопросом//
type Abonement struct {
	UID            UID    `json:"uid"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          string `json:"price"`
	ActivateTime   string `json:"activateTime"`
	//Места, где действует абонемент
	PlacesUIDs     []UID `json:"placesUIDs"`
	//Виды спорта относящиеся к абонементу
	CategoriesUIDs []UID `json:"categoriesUIDs"`
}
