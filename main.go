package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

var links = []string{}
var wg sync.WaitGroup

func crawlPage(url string) {
	fmt.Println("Crawling: ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	for {
		// Next token
		tt := z.Next()

		if tt == html.ErrorToken {
			return
		} else if tt == html.StartTagToken {
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				// Loop over anchor attributes and
				// if there is a 'href' then Print
				// the value.
				for _, a := range t.Attr {
					if a.Key == "href" {
						fmt.Println("Found link!: ", a.Val)
						links = append(links, a.Val)
						wg.Add(1)
						go crawlPage(a.Val)
						break
					}
				}
			}
		}
	}

	wg.Done()
}

func main() {

	resp, err := http.Get("https://www.fusion-conferences.com")
	if err != nil {
		log.Fatalln("Could not fetch URL")
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	for {
		// Next token
		tt := z.Next()

		if tt == html.ErrorToken {
			return
		} else if tt == html.StartTagToken {
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				// Loop over anchor attributes and
				// if there is a 'href' then Print
				// the value.
				for _, a := range t.Attr {
					if a.Key == "href" {
						fmt.Println("Found link!: ", a.Val)
						links = append(links, a.Val)
						wg.Add(1)
						go crawlPage(a.Val)
						break
					}
				}
			}
		}
	}

	wg.Wait()
}
