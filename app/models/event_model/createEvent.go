package event_model

import (
	"github.com/satori/go.uuid"
	"sport_kld/app/models"
)

func CreateEvent(event Event) (models.UID, error) {
	event.UID = models.UID(uuid.NewV1().String())
	if err := event.validate(); err != nil {
		return "", err
	}

	event.VisitorsNum = 0
	event.IsOver = false

	if err := putEvent(event); err != nil {
		return "", err
	}

	creatorRole := EventUserRole{
		UserUID:         event.CreatorUID,
		EventUID:        event.UID,
		Role:            "Организатор",
		RoleDescription: "Создатель события",
	}

	if err := putEventUserRole(creatorRole); err != nil {
		return "", err
	}

	return event.UID, nil
}
