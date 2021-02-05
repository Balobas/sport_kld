package main

import (
	"./models/place_model"
	"fmt"
	_ "fmt"
)

func main() {
	//p, err := place.GetPlacesByUIDs([]models.UID{})
	//fmt.Println(p, err)
	//
	//pl, er := place_controller.GetPlacesByTags("Пло")
	//fmt.Println(len(pl), er)

	//place := place_model.Place{}
	//num := reflect.TypeOf(&place).Elem().NumField()
	//
	//for i := 0; i < num; i++ {
	//	fmt.Println(reflect.TypeOf(&place).Elem().Field(i).Tag.Get("json"))
	//}

	p, e := place_model.GetPlacesByFields([]string{"webSite", "phones", "facebook"}, "ad")

	fmt.Println(p, e)

	//query := "INSERT INTO tags VALUES "
	//var inserts []string
	//var params []interface{}
	//
	//products := []models.Tag{
	//	{
	//		"adad",
	//		"adad",
	//	},
	//	{
	//		"okoko",
	//		"ononaoidn",
	//	},
	//}
	//
	//for _, v := range products {
	//	inserts = append(inserts, "(?, ?)")
	//	params = append(params, v.UID, v.Name)
	//}
	//queryVals := strings.Join(inserts, ",")
	//query = query + queryVals
	//log.Println("query is", query)
	//ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancelfunc()
	//stmt, err := database.MysqlDB.PrepareContext(ctx, query)
	//if err != nil {
	//	log.Printf("Error %s when preparing SQL statement", err)
	//}
	//defer stmt.Close()
	//
	//res, err := stmt.ExecContext(ctx, params...)
	//if err != nil {
	//	log.Printf("Error %s when inserting row into products table", err)
	//}
	//rows, err := res.RowsAffected()
	//if err != nil {
	//	log.Printf("Error %s when finding rows affected", err)
	//}
	//log.Printf("%d products created simultaneously", rows)
}

