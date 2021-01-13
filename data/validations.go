package data

func (org Organization) Validate() error {

	return nil
}

func (org Organization) IsValidUID() bool {
	return len(org.UID) != 0
}

func (place Place) Validate() error {

	return nil
}

func (place Place) IsValidUID() bool {
	return len(place.UID) != 0
}


func (abonement Abonement) Validate() error {

	return nil
}

func (abonement Abonement) IsValidUID() bool {
	return len(abonement.UID) != 0
}
