package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
)

func main() {

	resp, err := http.Get("https://www.fusion-conferences.com")
	if err != nil {
		log.Fatalln("Could not fetch URL")
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next();
		if tt == html.ErrorToken {
			return
		} else if tt == html.StartTagToken {
			t := z.Token()
			fmt.Println(t.String()+"\n")
		}
	}

}
