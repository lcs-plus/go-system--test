package mypack

import (
	_ "github.com/go-sql-driver/mysql"
)

var Mysql_host string = "127.0.0.1"
var Mysql_user string = "root"
var Mysql_pwd string = "root"
var Mysql_port string = "3306"
var Mysql_db string = "yingbei"
var Mysql_charset string = "utf8"

type Auth struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OrderXml struct {
	DataId             string
	OrderType          string
	OrderNo            string
	BatchNumbers       string
	EbpCode            string
	EbpName            string
	EbcCode            string
	EbcName            string
	BwsNo              string
	CustomsCode        int
	PortCode           int
	BuyerRegNo         string
	BuyerName          string
	BuyerIdNumber      string
	BuyerTelephone     string
	Consignee          string
	ConsigneeTelephone string
	ConsigneeAddress   string
	InsuredFee         int
	Freight            float32
	Discount           float32
	TaxTotal           float32
	Currency           int
	GrossWeight        float32
	NetWeight          float32
	Note               string
}

type OrderInfoClass struct {
	OrderId              string  `db:"order_id"`
	SubOrderNo           string  `db:"sub_order_no"`
	OuterSkuId           string  `db:"outer_sku_id"`
	CustomsCode          string  `db:"customs_code"`
	Num                  int     `db:"num"`
	FenxiaoPrice         float32 `db:"fenxiao_price"`
	CrossBorderTradeMode string  `db:"cross_border_trade_mode"`
	Payment              float32 `db:"payment"`
	DiscountPrice        float32 `db:"discount_price"`
	FenxiaoPayment       float32 `db:"fenxiao_payment"`
	TotalFee             float32 `db:"total_fee"`
	TaxTotal             float32 `db:"tax_total"`
	Freight              float32 `db:"freight"`
	FenxiaoTaxTotal      float32 `db:"fenxiao_tax_total"`
	FenxiaoFreight       float32 `db:"fenxiao_freight"`
	FenxiaoDiscount      float32 `db:"fenxiao_discount"`
	Discount             float32 `db:"discount"`

	Unit          int     `db:"unit"`
	RoughWeight   float32 `db:"rough_weight"`
	NetWeight     float32 `db:"net_weight"`
	Country       int     `db:"country"`
	Specification string  `db:"specification"`
	ShopName      string  `db:"shop_name"`
	Numed         string  `db:"numed"`
	ShopCode      string  `db:"shop_code"`
	United        string  `db:"united"`
	Numeder       string  `db:"numeder"`
	Uniteder      string  `db:"uniteder"`
}
