package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ConvertGolang(key string) string {
	splitted := strings.Split(key, "\\;")
	if len(splitted) == 0 {
		log.Fatalln("Error")
	}

	return splitted[0]

}
func CheckHTML(html string) {

	if apps, ok := apps["apps"].(map[string]interface{}); ok {
		for _, app := range apps {
			if appMap, ok := app.(map[string]interface{}); ok {
				if regexes, ok := appMap["html"].([]interface{}); ok {
					for _, regex := range regexes {
						if r, ok := regex.(string); ok {
							re := regexp.MustCompile(ConvertGolang(r))

							m := re.FindAllStringSubmatch(html, -1)
							if len(m) != 0 {
								fmt.Println(m)
							}
						}
					}
				}
			}
		}
	}
}

func CheckScripts(script string) {

	if apps, ok := apps["apps"].(map[string]interface{}); ok {
		for _, app := range apps {
			if appMap, ok := app.(map[string]interface{}); ok {
				if scriptSrc, ok := appMap["scriptSrc"].([]interface{}); ok {
					for _, src := range scriptSrc {
						if srcStr, ok := src.(string); ok {
							re, err := regexp.Compile(ConvertGolang(srcStr))
							if err != nil {
								fmt.Println(err)
							}
							m := re.FindAllStringSubmatch(script, -1)
							if len(m) != 0 {
								fmt.Println(m)
							}
						}
					}
				}
			}
		}
	}
}
