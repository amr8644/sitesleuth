package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Data struct {
	URL     string
	HTML    string
	Headers map[string][]string
	Cookies map[string]string
	Meta    map[string]string
	Script  map[string]string
}

var data Data

func ParseHTMLPage(page string) {
	regex_scripts := regexp.MustCompile(`<script.*?src="([^"]+)".*?</script>`)
	submatches := regex_scripts.FindAllStringSubmatch(data.HTML, -1)

	for _, match := range submatches {
		fmt.Println(match[1])
	}
    fmt.Println(data.HTML)
}

func ScrapeURL(response http.Response, value string) {
	headers := response.Header

	var ct string = response.Header.Get("Content-Type")

	if ct == "" || ct != "text/html" {
		log.Println("{} response use Content-Type {} but text/html is needed")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	data.URL = value
	data.HTML = string(body)

	for k, v := range headers {
		data.Headers = make(map[string][]string)
		data.Headers[k] = v
	}

	for _, cookie := range response.Cookies() {
		data.Cookies = make(map[string]string)
		data.Cookies[cookie.Name] = cookie.Value
	}

	ParseHTMLPage(string(body))
}

func SendRequests(value string) *http.Response {

	client := &http.Client{}
	request, err := http.NewRequest("GET", value, nil)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	return response
}
