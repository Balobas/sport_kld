package event_model

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"sport_kld/app/models"
	"sport_kld/database"
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

	event.VisitorsNum = 0
	event.IsOver = false

	result, err := database.MysqlDB.NamedExec("INSERT INTO events(uid, name, description, dates, time, visitorsNum, visitorsLimit, placeUid, creatorUid, isPrivate, isOver) VALUES (:uid, :name, :description, :dates, :time, :visitorsNum, :visitorsLimit, :placeUid, :creatorUid, :isPrivate, :isOver)", &event)
	if err != nil {
		return "", errors.New("cant create event")
	}

	// Заменить на логер
	a, _ := result.RowsAffected()
	fmt.Println("inserted ", a, " rows")

	return event.UID, nil
}
