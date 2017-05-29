package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getHTML(url string) (string, error) {
	tr := &http.Transport{DisableKeepAlives: false}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "Token"))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36")
	fmt.Println(req.Header)
	req.Close = false
	res, err := tr.RoundTrip(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	responseData, _ := ioutil.ReadAll(res.Body)
	responseString := string(responseData)
	if err != nil {
		log.Fatal(err)
	}
	b := []byte(responseString)
	r := MinifyHTML(b)
	return r, nil
}
