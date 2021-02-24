package event_model

import (
	"errors"
	"sport_kld/app/models"
)

func ChangeUserRole(role EventUserRole, executorUid models.UID) error {
	oldRole, err := GetEventUserRoleByUid(role.EventUID, role.UserUID)
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

	oldRole.Role = role.Role
	oldRole.RoleDescription = role.RoleDescription

	return putEventUserRole(oldRole)
}
