package comments

import (
	"sport_kld/database"
)

func putComment(comment Comment) error {
	_, err := database.MysqlDB.NamedExec("insert into comments(uid, text, authorUid, time) values (:uid, :text, :authorUid, :time)", &comment)
	return err
}
