package mypack

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var mysql_host string = "127.0.0.1"
var mysql_user string = "root"
var mysql_pwd string = "root"
var mysql_port int = 3306



type Auth struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OrderXml struct {
	dataId             string
	orderType          string
	orderNo            string
	batchNumbers       string
	ebpCode            string
	ebpName            string
	ebcCode            string
	ebcName            string
	bwsNo              string
	customsCode        int
	portCode           int
	buyerRegNo         string
	buyerName          string
	buyerIdNumber      string
	buyerTelephone     string
	consignee          string
	consigneeTelephone string
	consigneeAddress   string
	insuredFee         int
	freight            float32
	discount           float32
	taxTotal           float32
	currency           int
	grossWeight        float32
	netWeight          float32
	note               string
}

func CheckError(err error) {

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

}

func MySqlQueryOne()  {



}