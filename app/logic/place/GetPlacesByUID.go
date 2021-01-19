package place

import (
	"../../../core/actions"
	"../../../data"
	"fmt"
	"github.com/pkg/errors"
)

func GetPlaceByUID(uid data.UID) (data.Place, error) {
	var place data.Place
	if err := actions.ActionGet(uid, &place); err != nil {
		return data.Place{}, errors.WithStack(err)
	}
	fmt.Println(place)
	return place, nil
}

func GetPlacesByUIDs(uids []data.UID) ([]data.UID, error) {
	return nil, nil
}