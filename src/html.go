package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getHTML(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	if err != nil {
		log.Fatal(err)
	}
	b := []byte(responseString)
	r := MinifyHTML(b)
	return r, nil
}
