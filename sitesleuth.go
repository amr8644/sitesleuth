package main

import "strings"

func Start() {
	//tar := "https://www.shielder.com"

	tar := "https://www.w3schools.com/"
	//ParseResponse(*response)
	links := LinkGrabber(tar)
	data.URL = tar
	for _, v := range links {
		if v[0] == '/' {
			//		fmt.Println(v)
		}
		if strings.Contains(v, tar) {
			ScrapeURL(tar)
		}
	}

    CheckHTML(data.HTML)

}
