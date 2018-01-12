package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hey")
	resp, err := http.Get("https://www.fusion-conferences.com")
	if err != nil {
		log.Fatalln("Could not fetch URL")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Could not read body")
	}
	fmt.Println(string(body))
}
