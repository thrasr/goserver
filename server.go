package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", html.EscapeString(r.URL.Path))
}

func main() {

	//handler functions
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world")
	})
	http.HandleFunc("/", echo)

	//grab openshift's IP and port number
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	//run server
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		//panic(err)       //panic shuts down server
		log.Fatal(err)
	}

}
