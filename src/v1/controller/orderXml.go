package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"v1/mypack"
)

func GetOrderXml(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "*")

	var tids = r.PostFormValue("tid")

	tids = strings.Replace(tids, "undefined,", "", -1)

	tids = tids[:len(tids)-1]

	tids_list := strings.Split(tids, ",")

	db, err := sql.Open("mysql", mypack.Mysql_user+":"+mypack.Mysql_pwd+"@tcp("+mypack.Mysql_host+":"+mypack.Mysql_port+")/"+mypack.Mysql_db+"?charset="+mypack.Mysql_charset)

	mypack.CheckError(err)

	for _, tid := range tids_list {

		sqls := "SELECT `order_id`,`sub_order_no`,`outer_sku_id`,`customs_code`,`num`,`fenxiao_price`,`cross_border_trade_mode`,`payment`,`discount_price`,`fenxiao_payment`,`total_fee`,`tax_total`,`freight`,`fenxiao_tax_total`,`fenxiao_freight`,`fenxiao_discount`,`discount`,`unit`,`rough_weight`,`net_weight`,`country`,`specification`,`shop_name`,`numed`,`shop_code`,`united`,`numeder`,`uniteder` FROM `sx_order_shop`as shop INNER JOIN `sx_order_code` AS code ON shop.`outer_sku_id`=code.`key` WHERE shop.`order_id` =" + tid

		mypack.CheckError(err)

		rows, err := db.Query(sqls)

		mypack.CheckError(err)

		var orderShopList []*mypack.OrderInfoClass

		for rows.Next() {

			var info mypack.OrderInfoClass

			err := rows.Scan(&info.OrderId, &info.SubOrderNo, &info.OuterSkuId, &info.CustomsCode, &info.Num, &info.FenxiaoPrice, &info.CrossBorderTradeMode, &info.Payment, &info.DiscountPrice, &info.FenxiaoPayment, &info.TotalFee, &info.TaxTotal, &info.Freight, &info.FenxiaoTaxTotal, &info.FenxiaoFreight, &info.FenxiaoDiscount, &info.Discount, &info.Unit, &info.RoughWeight, &info.NetWeight, &info.Country, &info.Specification, &info.ShopName, &info.Numed, &info.ShopCode, &info.United, &info.Numeder, &info.Uniteder)

			mypack.CheckError(err)

			orderShopList = append(orderShopList, &info)

		}

		for _, orderInfo := range orderShopList {

			fmt.Println(orderInfo)

		}

		var DataHead mypack.OrderXml

		DataHead.DataId = "MA6XNJ0H4CEB311" + time.Now().Format("20060102") + mypack.GetRandomString(7)
		DataHead.OrderType = "I"
		DataHead.OrderNo = orderShopList[0].SubOrderNo
		DataHead.BatchNumbers = ""
		DataHead.EbpCode = ""
		DataHead.EbpName = ""
		DataHead.EbcCode = ""
		DataHead.EbcName = ""

	}

	var result mypack.Resp

	result.Code = 200
	result.Msg = "ok"

	json.NewEncoder(w).Encode(result)

}

/*
	DataHead['dataId'] = ""
	DataHead['orderType'] = ""
	DataHead['orderNo'] = ""
	DataHead['batchNumbers'] = ""
	DataHead['ebpCode'] = ""
	DataHead['ebpName'] = ""
	DataHead['ebcCode'] = ""
	DataHead['ebcName'] = ""
	DataHead['bwsNo'] = ""
	DataHead['customsCode'] = ""
	DataHead['portCode'] = ""
	DataHead['buyerRegNo'] = ""
	DataHead['buyerName'] = ""
	DataHead['buyerIdNumber'] = ""
	DataHead['buyerTelephone'] = ""
	DataHead['consignee'] = ""
	DataHead['consigneeTelephone'] = ""
	DataHead['consigneeAddress'] = ""
	DataHead['insuredFee'] = ""
	DataHead['freight'] = ""
	DataHead['discount'] = ""
	DataHead['taxTotal'] = ""
	DataHead['currency'] = ""
	DataHead['grossWeight'] = ""
	DataHead['netWeight'] = ""
	DataHead['note'] = ""
*/
