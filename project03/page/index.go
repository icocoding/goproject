package project03

import (
	"fmt"
	"net/http"
)

func IndexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, index page")
}
