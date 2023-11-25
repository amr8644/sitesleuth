package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Data struct {
	URL     string
	HTML    string
	Headers map[string][]string
	Cookies map[string]string
	Meta    []string
	Script  []string
}

var data Data

func PrintData() {

	fmt.Println("URL:", data.URL)


	fmt.Println("Headers:")
	for k, v := range data.Headers {

		fmt.Println("- ", k, ":", v)
	}

	fmt.Println("Cookies:")
	for k, v := range data.Cookies {
		fmt.Println("- ", k, ":", v)
	}

	fmt.Println("Meta:")
	for k, v := range data.Meta {
		fmt.Println("- ", k, ":", v)
	}

	fmt.Println("Script:")
	for k, v := range data.Script {
		fmt.Println("- ", k, ":", v)
	}


	//fmt.Println("HTML:", data.HTML)

}

func LinkGrabber(value string) []string {

	var links []string
	res := SendRequests(value)
	body, err := io.ReadAll(res.Body)

	doc, err := html.Parse(strings.NewReader(string(body[:])))

	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return links
}

func ParseHTMLPage(page string) {
	// Parse HTML
	doc, err := html.Parse(strings.NewReader(data.HTML))

	if err != nil {
		panic(err)
	}
	// Find nodes
	var meta, scripts []string

	IterateNodes(doc, func(n *html.Node) {
		if n.Data == "meta" {
			meta = append(meta, GetAttr(n, "content"))
		}
		if n.Data == "script" {
			scripts = append(scripts, GetAttr(n, "src"))
		}
	})

	data.Meta = meta
	data.Script = scripts
	// Print results
}

func ParseResponse(response http.Response) {
	headers := response.Header
	for k, v := range headers {
		if k == "Set-Cookie" {
			data.Cookies = make(map[string]string)
			data.Cookies[k] = v[0]
		} else {
			data.Headers = make(map[string][]string)
			data.Headers[k] = v
		}

	}
}

func ParseRequest(request http.Request) {
	headers := request.Header
	for k, v := range headers {
		if k == "Set-Cookie" {
			data.Cookies = make(map[string]string)
			data.Cookies[k] = v[0]
		} else {
			data.Headers = make(map[string][]string)
			data.Headers[k] = v
		}
	}
}

func ScrapeURL(value string) (error,http.Response) {
	response := SendRequests(value)
	//headers := response.Header

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//data.URL = value
	data.HTML = string(body)

	ParseResponse(*response)
	ParseHTMLPage(string(body))

	return nil,*response
}

func SendRequests(value string) *http.Response {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", value, nil)
	request.Header.Set("User-Agent", RandomUserAgents())

	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	return response
}
