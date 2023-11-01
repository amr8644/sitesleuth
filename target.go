package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func PrintData() {

	fmt.Println("URL:", data.URL)

	//fmt.Println("HTML:", data.HTML)

	fmt.Println("Headers:")
	for k, v := range data.Headers {
		if k == "X-Powered-By" {

			fmt.Println("- ", k, ":", v)
		}
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

}

func LinkGrabber(value string) []string {

    var links []string
		res := SendRequests(value)
	body, err := ioutil.ReadAll(res.Body)

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
	// Print results
	//fmt.Println("Meta:", meta)
	//fmt.Println("Scripts:", scripts)
}

func ParseResponse(response http.Response) {
	headers := response.Header
	for k, v := range headers {
		if k == "Set-Cookie" {
			data.Cookies = make(map[string]string)
			data.Cookies[k] = v[0]
			//fmt.Println(k, v)
		} else {
			data.Headers = make(map[string][]string)
			data.Headers[k] = v
		}

	}
}

func ParseRequest(response http.Response) {

	headers := response.Header

	for k, v := range headers {
		if k == "Set-Cookie" {
			data.Cookies = make(map[string]string)
			data.Cookies[k] = v[0]
			fmt.Println(k, v)
		} else {
			data.Headers = make(map[string][]string)
			data.Headers[k] = v
		}
	}
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
        fmt.Println(k,v)
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
