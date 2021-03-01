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

	creatorEventUser := EventUser{
		UserUID:         event.CreatorUID,
		EventUID:        event.UID,
		Role:            "Организатор",
		RoleDescription: "Создатель события",
	}

	if err := putEventUser(creatorEventUser); err != nil {
		return "", err
	}

	//if _, err := database.MysqlDB.Exec("INSERT INTO event_users(event_uid, user_uid) VALUES (?, ?)", event.UID, event.CreatorUID); err != nil {
	//	return "", errors.New("cant join to event")
	//}
	return event.UID, nil
}
