package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"v1/mypack"
)

func GetOrderList(writer http.ResponseWriter, _ *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "*")

	var sql_url = "root:root@/test?charset=utf8"

	db, err := sql.Open("mysql", sql_url)

	mypack.CheckError(err)

	defer db.Close()

	var sqls = "SELECT `tid`,`id`,`id_card_name`,`id_card_number` FROM `sx_order_order` WHERE `is_del`=1 limit 100000"

	rows, err := db.Query(sqls)

	mypack.CheckError(err)

	type orderSql struct {
		Tid          string `db:"tid"`
		Id           int    `db:"id"`
		IdCardName   string `db:"id_card_name"`
		IdCardNumber string `db:"id_card_number"`
	}

	var res []*orderSql

	for rows.Next() {

		var info orderSql

		err := rows.Scan(&info.Tid,&info.Id, &info.IdCardName, &info.IdCardNumber)

		mypack.CheckError(err)

		res = append(res, &info)

	}

	err = json.NewEncoder(writer).Encode(res)

	mypack.CheckError(err)

}