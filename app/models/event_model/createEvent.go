package event_model

import (
	"errors"
	"github.com/satori/go.uuid"
	"sport_kld/app/models"
)

func CreateEvent(event Event) (models.UID, error) {
	event.UID = models.UID(uuid.NewV1().String())

	if len(event.Name) == 0 {
		return "", errors.New("empty event name field")
	}
	if len(event.Description) == 0 {
		return "", errors.New("empty event description field")
	}
	if len(event.Dates) == 0 {
		return "", errors.New("empty event dates field")
	}
	if len(event.Time) == 0 {
		return "", errors.New("empty event time field")
	}
	if event.VisitorsLimit == 0 {
		return "", errors.New("empty event visitors limit field")
	}
	if len(event.PlaceUID) == 0 {
		return "", errors.New("empty event place uid field")
	}
	if len(event.CreatorUID) == 0 {
		return "", errors.New("empty event creator uid field")
	}
	if event.IsPrivate && len(event.EventPassword) == 0 {
		return "", errors.New("if you create private event, you should enter password")
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
