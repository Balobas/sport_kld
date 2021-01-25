package place

import (
	"../../../core/actions"
	"../../../models"
	"fmt"
	"github.com/pkg/errors"
)

func GetPlaceByUID(uid models.UID) (models.Place, error) {
	var place models.Place
	if err := actions.ActionGet(uid, &place); err != nil {
		return models.Place{}, errors.WithStack(err)
	}
	fmt.Println(place)
	return place, nil
}

func GetPlacesByUIDs(uids []models.UID) ([]models.UID, error) {
	return nil, nil
}
