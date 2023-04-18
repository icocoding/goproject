package main

import (
	"fmt"
	"encoding/json"
	"github.com/tidwall/gjson"
)
/**
* 从网络获取数据, 并解析为json 结构化
*/


type Group struct {
	Id string `json:"id"`
	Name string `json:"name"`
	GroupNumber string `json:"group_number"`
	GroupKey string `json:"group_key"`
}
type MyData struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Icon   string   `json:"icon"`
	Status string   `json:"status"`
	Groups []Group `json:"groups"`
}
type HttpResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data struct {
        List []MyData `json:"list"`
    } `json:"data"`
}
func Test02()  {
	data := Test01()
	fmt.Println("======> get data from test01")

	var respData HttpResp
	if err := json.Unmarshal([]byte(data), &respData); err != nil {
		fmt.Println("parse err", err)
		return
	}
	fmt.Println("======> use encoding/json")
	fmt.Println(respData.Code)
	fmt.Println(respData.Msg)

	var myData []MyData
	myData = respData.Data.List

	fmt.Println(myData[0].ID, myData[0].Name)

	fmt.Println("======> use gson")
	m, ok := gjson.Parse(data).Value().(map[string]interface{})
	if !ok {
		// not a map
	}
	fmt.Println(m["code"])
	fmt.Println(m["msg"])
	// fmt.Println(m["data"])

	a := gjson.Get(data, "data.list.0")
	fmt.Println(a)
}