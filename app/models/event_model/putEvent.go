package event_model

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func putEvent(event Event) error {
	if _, err := GetEventByUid(event.UID); err == nil {
		_, err = database.MysqlDB.NamedExec(
			"update events set name = :name, description = :description, dates = :dates, time = :time," +
			"visitorsNum = :visitorsNum, visitorsLimit = :visitorsLimit, placeUid = :placeUid, creatorUid = :creatorUid, isPrivate = :isPrivate, isOver = :isOver" +
			" where uid = :uid", &event)

		return err
	}

	result, err := database.MysqlDB.NamedExec("INSERT INTO events(uid, name, description, dates, time, visitorsNum, visitorsLimit, placeUid, creatorUid, isPrivate, isOver) VALUES (:uid, :name, :description, :dates, :time, :visitorsNum, :visitorsLimit, :placeUid, :creatorUid, :isPrivate, :isOver)", &event)
	if err != nil {
		return errors.New("cant put event")
	}

	// Заменить на логер
	a, _ := result.RowsAffected()
	fmt.Println("inserted ", a, " rows")

	return nil
}
