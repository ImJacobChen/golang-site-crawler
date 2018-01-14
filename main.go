package main

import (
	"fmt"
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
			break
		}

		// if tt == html.StartTagToken {
		// 	t := z.Token()

		// 	isAnchor := t.Data == "a"
		// 	if isAnchor {
		// 		// Loop over anchor attributes and
		// 		// if there is a 'href' then Print
		// 		// the value.
		// 		for _, a := range t.Attr {
		// 			if a.Key == "href" {
		// 				fmt.Println("Found link!: ", a.Val)
		// 				fmt.Println(z.Text())
		// 				links = append(links, a.Val)
		// 				wg.Add(1)
		// 				go crawlPage(a.Val)
		// 				break
		// 			}
		// 		}
		// 	}
		// }

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

		// if tt == html.TextToken {
		// 	token := z.Token()

		// 	fmt.Println("Title = ", token.Data)
		// }
	}

	//wg.Done()
}

func main() {

	crawlPage("https://www.fusion-conferences.com")

	//wg.Wait()
}
