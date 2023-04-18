package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)
/**
* 从网络获取数据, 并解析为json(非结构化)
*/
func Test01() string  {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.bejson.com/Bejson/Api/LanguageGroup/getGroupList?site_id=2", nil)
	if err != nil {
		fmt.Println("req init err:", err)
		return ""
	}
	
	req.Header.Set("referer", "https://www.bejson.com")
	req.Header.Set("orgin", "https://www.bejson.com")
	req.Header.Set("accept", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("req get err:", err)
		return ""
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return ""
	}
	fmt.Println(string(body))

	//数组解析
	// var p []interface{}
	// err = json.Unmarshal([]byte(body), &p)
	//map解析
	var p interface{}
	err = json.Unmarshal([]byte(body), &p)
	if err != nil {
		fmt.Println("parse body err:", err)
		return ""
	}
	m := p.(map[string]interface{}) 

	fmt.Println("=====> resp msg", m["msg"])

	for k, v := range m {
		fmt.Println(k, "=", v)
	}
	return string(body)
}