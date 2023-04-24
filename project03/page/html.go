package project03

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HtmlPage(res http.ResponseWriter, req *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("cwd error")
		panic(err)
		// return
	}
	fmt.Println("current dir: ", dir)
	// /Users/etoneyang/Documents/Project/goproject/project03/html/index.html
	content, _ := ioutil.ReadFile("./project03/html/index.html")
	len, err := res.Write(content)
	if err != nil {
		fmt.Println("Write error", err)
		return
	}
	fmt.Println(len)
}
