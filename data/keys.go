package data

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"../core"
)

//Organization
func (org Organization) GenerateAndSetKey() (core.Object, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	org.UID = UID(uid.String())
	return org, nil
}

func (org Organization) GetKey() string {
	return org.UID.String()
}

func (org Organization) SetKey(key string) (core.Object, error) {
	if !(UID(key)).isCorrect() {
		return nil, errors.Errorf("uid is not correct %s", key)
	}
	org.UID = UID(key)
	return org, nil
}

//Place
func (place Place) GenerateAndSetKey() (core.Object, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	place.UID = UID(uid.String())
	return place, nil
}

func (place Place) GetKey() string {
	return place.UID.String()
}

func (place Place) SetKey(key string) (core.Object, error) {
	if !(UID(key)).isCorrect() {
		return nil, errors.Errorf("uid is not correct %s", key)
	}
	place.UID = UID(key)
	return place, nil
}

//Abonement
func (abonement Abonement) GenerateAndSetKey() (core.Object, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	abonement.UID = UID(uid.String())
	return abonement, nil
}

func (abonement Abonement) GetKey() string {
	return abonement.UID.String()
}

func (abonement Abonement) SetKey(key string) (core.Object, error) {
	if !(UID(key)).isCorrect() {
		return nil, errors.Errorf("uid is not correct %s", key)
	}
	abonement.UID = UID(key)
	return abonement, nil
}
