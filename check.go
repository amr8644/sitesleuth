package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"
)

var result map[string]interface{}

func ConvertGolang(key string) string {

	escapedValue := regexp.QuoteMeta(key)
	htmlPattern := `(?s)` + escapedValue

	return htmlPattern
}

func CheckHTML(html string) {

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
						regexes := ConvertGolang(str)
						matched, err := regexp.MatchString(regexes, html)
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						if matched {
							fmt.Println(d["name"])
						}
					}
				}
			}
		}
	}
}

func CheckScripts(scripts []string) {

    fmt.Println(scripts)
}

func ReadFile() {

    // Record the start time
	startTime := time.Now()
	// Record the end time
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

	endTime := time.Now()
	// Calculate the duration
	duration := endTime.Sub(startTime)

	// Print the duration
	fmt.Printf("Time taken: %v\n", duration)


}
