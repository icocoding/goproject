package project02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

/*
读取文件内容
*/
func Test01() {
	path := "./project02/files/test.json"
	fmt.Println("project 02 test 01: File read & write")
	fileData, err := ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	id := GetId()
	var j interface{}
	err = json.Unmarshal(fileData, &j)
	if err != nil {
		fmt.Println("parse json err:", err)
		return
	}
	m := j.(map[string]interface{})
	oldId := m["id"]
	//赋值给id
	m["id"] = string(id)
	fmt.Println(oldId, m)
	str, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("Marshal json err:", err)
		return
	}
	//写入文件
	err = os.WriteFile(path, str, os.ModePerm)
	if err != nil {
		fmt.Println("WriteFile err:", err)
		return
	}
	fmt.Println("WriteFile success")
}
func ReadFile(path string) ([]byte, error) {
	dir, err := os.Getwd()
	b0 := make([]byte, 0)
	if err != nil {
		fmt.Println("cwd error")
		return b0, err
	}
	fmt.Println("current dir: ", dir)
	// 打开文件
	jsonFile, err := os.Open(path)

	if os.IsExist(err) {
		fmt.Println("open test.json file error")
		return b0, err
	}
	defer jsonFile.Close()

	// 读取文件内容
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("read test.json file error")
		return b0, err
	}

	fmt.Println(string(jsonData))
	return jsonData, nil
}

/*
时间格式化为Id
*/
func GetId() string {
	timeNow := time.Now()

	timeMillis := timeNow.UTC()

	fmt.Println(timeMillis)

	//string 格式, 其中"2006-01-02 15:04:05"是系统规定的模版，值不能变
	fmt.Println(timeNow.Format("2006-01-02 15:04:05"))

	id := timeNow.Format("20060102150405")
	fmt.Println(id)

	return id
}
