package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 0600 is the same as chmod 600
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func main() {
	// Think routes.  I say think, but that's exactly what it is.
	http.HandleFunc("/view/", viewHandler)
	// That's it?!
	// Literally all we need for async is the go keyword.
	go testAPICall()
	fmt.Println("Starting server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Will NEVER get past this line before termination.
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func testAPICall() {
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("API call failed with", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
