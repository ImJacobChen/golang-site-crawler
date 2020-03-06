package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var links = []string{}
var crawledPages = []string{}
var wg sync.WaitGroup
var startURL *url.URL

func crawlPage(url string) {
	fmt.Println("Crawling: ", url)
	crawledPages = append(crawledPages, url)
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
			break
		}

		if tt == html.StartTagToken {
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
						if pageShouldBeCrawled(a.Val) {
							wg.Add(1)
							go crawlPage(a.Val)
						}
						break
					}
				}
			}
		}

		if tt == html.StartTagToken {
			t := z.Token()
		 	if t.Data == "title" {
		 		tokenType := z.Next()
		 		if tokenType == html.TextToken {
		 			titleToken := z.Token()
		 			fmt.Println("Found title: ", titleToken.String())
		 		}

		 	}

		}

		if tt == html.TextToken {
	 		token := z.Token()

		 	fmt.Println("Title = ", token.Data)
		}
	}

	wg.Done()
}

func pageShouldBeCrawled(url string) bool {

	crawledPagesLen := len(crawledPages)
	for i := 0; i < crawledPagesLen; i++ {
		if url == crawledPages[i] {
			return false
		}
	}

	if !strings.Contains(url, startURL.Hostname()) {
		return false
	}

	return true
}

func main() {

	url, err := url.Parse("https://www.fusion-conferences.com")
	if err != nil {
		log.Fatalln("Invalid start URL")
	}

	startURL = url
	wg.Add(1)
	crawlPage(startURL.String())

	wg.Wait()
}
