package main

import (
	"fmt"
	"io/ioutil"
	"log"
    "golang.org/x/net/html"
  "strings"
	"net/http"
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

func getAttr(n *html.Node, attr string) string {
  for _, a := range n.Attr {
    if a.Key == attr {
      return a.Val
    }
  }
  return ""
}

// Iterate DOM nodes
func iterateNodes(n *html.Node, f func(*html.Node)) {
  f(n)
  for c := n.FirstChild; c != nil; c = c.NextSibling {
    iterateNodes(c, f)
  }
}

func ParseHTMLPage(page string) {


	// Parse HTML
	doc, err := html.Parse(strings.NewReader(data.HTML))
	if err != nil {
		panic(err)
	}

	// Find nodes
	var meta, scripts []string

	iterateNodes(doc, func(n *html.Node) {
		if n.Data == "meta" {
			meta = append(meta, getAttr(n, "content"))
		}
		if n.Data == "script" {
			scripts = append(scripts, getAttr(n, "src"))
		}
	})

	// Print results
	fmt.Println("Meta:", meta)
	fmt.Println("Scripts:", scripts)

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
