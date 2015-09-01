package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	//"os"     //command line argument to pass in port?
)

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", html.EscapeString(r.URL.Path))
}

func main() {

	//var h Hello
	//var e Echo

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world")
	})
	http.HandleFunc("/", echo)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		//panic(err)       //panic shuts down server
		log.Fatal(err)
	}

}
