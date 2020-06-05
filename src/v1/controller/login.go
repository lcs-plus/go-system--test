package controller

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"v1/mypack"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Access-Control-Allow-Origin","*")
	writer.Header().Set("Access-Control-Allow-Headers","*")

	/*_ = request.ParseMultipartForm(128)

	form := request.Form

	fmt.Println(form)

	fmt.Println(request)*/

	var auth mypack.Auth

	err := json.NewDecoder(request.Body).Decode(&auth)

	mypack.CheckError(err)

	var sql_url = "root:root@/test?charset=utf8"

	db,err := sql.Open("mysql",sql_url)

	mypack.CheckError(err)

	defer db.Close()

	var sql = "select `id`,`name`,`pwd` from `test_name` where `name`='" + auth.Username + "'"

	rows,err := db.Query(sql)

	mypack.CheckError(err)

	type userMql struct {
		Id    int  	  `db:"id"`
		Name  string  `db:"name"`
		Pwd   string  `db:"pwd"`
	}

	var info userMql

	for rows.Next() {

		err = rows.Scan(&info.Id,&info.Name,&info.Pwd)

		mypack.CheckError(err)

	}

	var result mypack.Resp

	if auth.Username == info.Name && auth.Pwd == info.Pwd {
		result.Code = 2000
		result.Msg  ="登录成功"
	}else{
		result.Code = 4000
		result.Msg  = "登录失败,用户名或密码错误"
	}

	err = json.NewEncoder(writer).Encode(result)

	mypack.CheckError(err)

}
