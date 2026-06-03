package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ggicci/httpin"
)

func main() {
	payloadStruct()
}

func urlencoded() {
	s := http.Server{Addr: "127.0.0.1:8888"}
	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("content-type", r.Header.Get("content-type"))
		//fmt.Println("parse", r.ParseForm())
		fmt.Println("form", r.Form)
		fmt.Println("form.test", r.Form.Get("test"))
		b, err := io.ReadAll(r.Body)
		fmt.Println("body", err, string(b))
		v, err := url.ParseQuery(string(b))
		fmt.Println("parsed query", err, v)
	})
	s.ListenAndServe()
}

func payloadStruct() {
	type Test struct {
		Path          string `in:"path=id"`
		First         string `in:"nonzero;form=first;"`
		Second        int    `in:"form=second;omitempty"`
		SomethingElse string `in:"form=something_else;omitempty"`
	}

	r, err := httpin.NewRequest("POST", "https://localhost/{id}", &Test{Path: "hello", First: "hello"})
	fmt.Println(err, r.URL.Path)

}
