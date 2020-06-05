package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"v1/mypack"
)

func GetOrderXml(w http.ResponseWriter,r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Headers","*")

	tids := r.PostFormValue("tid")

	appkey := r.PostFormValue("appkey")

	token := r.PostFormValue("token")

	fmt.Println("tids:" + tids)
	fmt.Println("appkey:" + appkey)
	fmt.Println("token:" + token)


	var xml mypack.OrderXml

	var result mypack.Resp

	result.Code = 200
	result.Msg = "ok"

	json.NewEncoder(w).Encode(result)



}