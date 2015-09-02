package main

import (
	"fmt"
	"github.com/chimeracoder/anaconda"
	"html"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

// grab and return a tweet with given hashtag
func hashtag(w http.ResponseWriter, r *http.Request) {
	// DONT UPLOAD KEY/SECRET TOKENS ANYWHERE!!!
	anaconda.SetConsumerKey("INSERTKEYHERE")
	anaconda.SetConsumerSecret("INSERTSECRETHERE")
	api := anaconda.NewTwitterApi("ACCESSTOKEN", "ACCESSTOKENSECRET")

	// grab small amount of tweets for testing
	v := url.Values{}
	v.Set("count", "10")

	// send search query with additional values v
	search_result, err := api.GetSearch("#"+r.URL.Path[1:], v)

	if err != nil {
		log.Fatal(err)
	}

	//pick a random tweet from results returned
	//may need an additional check here in case no tweets returned
	tweet := search_result.Statuses[rand.Intn(len(search_result.Statuses))]
	fmt.Fprint(w, tweet.User.Name+" @"+tweet.User.ScreenName+"\n")
	fmt.Fprint(w, tweet.Text)
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", html.EscapeString(r.URL.Path)[1:])
}

func main() {

	//handler functions

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world")
	})
	//http.HandleFunc("/", echo)
	http.HandleFunc("/", hashtag)

	//grab openshift's IP and port number
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	//run server
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		fmt.Println("ERROR")
		//panic(err)       //panic shuts down server
		log.Fatal(err)
	}

}
