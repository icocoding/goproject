package project03

import (
	project03 "cocoding-org/goproject/project03/page"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", project03.IndexPage)
	http.HandleFunc("/hello", project03.HelloPage)
	http.HandleFunc("/html", project03.HtmlPage)
	http.HandleFunc("/post", project03.PostAction)
	log.Println("start server on port:", 8099)
	log.Fatal(http.ListenAndServe("localhost:8099", nil))
}
