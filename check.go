package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

var result map[string]interface{}

func ConvertGolang(key string) string {

	escapedValue := regexp.QuoteMeta(key)
	htmlPattern := `(?s)` + escapedValue
	fmt.Println(htmlPattern+"\n")

	return htmlPattern
}

func CheckHTML(html string) []string {

    var regexes []string
	for name, info := range result {
		data, ok := info.(map[string]interface{})
		if !ok {
			fmt.Println("Invalid data format for app:", name)
			continue
		}
		for _, v := range data {
			d := v.(map[string]interface{})
			if array, ok := d["html"].([]interface{}); ok {
				for _, value := range array {
					if str, ok := value.(string); ok {
                       regexes =append(regexes, ConvertGolang(str))
					}
				}
			}
		}
	}
    return regexes
}


func ReadFile() {

	file, err := os.Open("apps.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	// Unmarshal the JSON data into a map
	err = json.Unmarshal(content, &result)
	if err != nil {
		panic(err)
	}

}
