package event_model

import (
	"errors"
	"sport_kld/app/models"
)

func ChangeUserRole(role EventUser, executorUid models.UID) error {
	oldEventUser, err := GetEventUserByUid(role.EventUID, role.UserUID)
	if err != nil {
		return err
	}

	event, err := GetEventByUid(role.EventUID)
	if err != nil {
		return err
	}

	if event.CreatorUID != executorUid {
		return errors.New("only creator can change roles")
	}

	if event.IsOver {
		return errors.New("event is over")
	}

	oldEventUser.Role = role.Role
	oldEventUser.RoleDescription = role.RoleDescription

	return putEventUser(oldEventUser)
}
