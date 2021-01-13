package data

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type OpeningHours struct {
	OpenNow     bool     `json:"open_now"`
	WeekdayText []string `json:"weekday_text"`
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

type LocationInfo struct {
	Address      string       `json:"adress"`
	Coordinates  Coordinates  `json:"location"`
	OpeningHours OpeningHours `json:"opening_hours"`
}

type Organization struct {
	UID            UID         `json:"uid"`
	Name           string      `json:"name"`
	MainAddress    string      `json:"mainAddress"`
	Abonements     []Abonement `json:"abonements"`
	Places         []Place     `json:"places"`
	CategoriesUIDs []UID       `json:"categoriesUIDs"`
	TagsUIDs       []UID       `json:"tagsUIDs"`
}

type Place struct {
	UID                UID            `json:"uid"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	HolderOrganization Organization   `json:"organization"`
	BasedOrganizations []Organization `json:"basedOrganizations"`
	FreeVisit          bool           `json:"freeVisit"`
	Abonements         []Abonement    `json:"abonements"`
	Location           LocationInfo   `json:"location"`
	CategoriesUIDs     []UID          `json:"categoriesUIDs"`
	TagsUIDs           []UID          `json:"tagsUIDs"`
}

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
