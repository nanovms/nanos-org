package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	index, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	sindex := string(index)

	book, err := ioutil.ReadFile("thebook.html")
	if err != nil {
		fmt.Println(err)
	}
	sbook := string(book)

	faq, err := ioutil.ReadFile("faq.html")
	if err != nil {
		fmt.Println(err)
	}
	sfaq := string(faq)

	community, err := ioutil.ReadFile("community.html")
	if err != nil {
		fmt.Println(err)
	}
	scommunity := string(community)

	terms, err := ioutil.ReadFile("terms.html")
	if err != nil {
		fmt.Println(err)
	}
	sterms := string(terms)

	privacy, err := ioutil.ReadFile("privacy.html")
	if err != nil {
		fmt.Println(err)
	}
	sprivacy := string(privacy)

	gettingstarted, err := ioutil.ReadFile("getting_started.html")
	if err != nil {
		fmt.Println(err)
	}
	sgettingstarted := string(gettingstarted)

	http.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, sfaq)
	})

	http.HandleFunc("/community", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, scommunity)
	})

	http.HandleFunc("/thebook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, sbook)
	})

	http.HandleFunc("/getting_started", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, sgettingstarted)
	})

	http.HandleFunc("/terms", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, sterms)
	})

	http.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, sprivacy)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, sindex)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			fmt.Println("ListenAndServe error: %v", err)
		}

	}()

	err = http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}
