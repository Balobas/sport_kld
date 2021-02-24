package event_model

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func putEvent(event Event) error {
	result, err := database.MysqlDB.NamedExec("INSERT INTO events(uid, name, description, dates, time, visitorsNum, visitorsLimit, placeUid, creatorUid, isPrivate, isOver) VALUES (:uid, :name, :description, :dates, :time, :visitorsNum, :visitorsLimit, :placeUid, :creatorUid, :isPrivate, :isOver)", &event)
	if err != nil {
		return errors.New("cant put event")
	}

	// Заменить на логер
	a, _ := result.RowsAffected()
	fmt.Println("inserted ", a, " rows")

	return nil
}
