package main

import (
	"net/http"
	"v1/controller"
)

func main() {

	/*
		//登录
		http.HandleFunc("/login",controller.Login)

		//注册
		http.HandleFunc("/get-order-list",controller.GetOrderList)
	*/

	http.HandleFunc("/get-order-xml",controller.GetOrderXml)

	http.ListenAndServe("0.0.0.0:8080", nil)

}
