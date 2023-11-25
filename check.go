package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

var result map[string]map[string]interface{}

func ConvertGolang(key string) string {

	escapedValue := regexp.QuoteMeta(key)
	htmlPattern := `(?s)` + escapedValue

	return htmlPattern
}

func CheckHTML(html string) {

}


func CheckScripts(scripts []string) {
	for _, appData := range result["apps"] {
		if appDataMap, ok := appData.(map[string]interface{}); ok {
			if scriptValue, ok := appDataMap["script"].(string); ok {

			} else {
				continue
			}
		} else {
			fmt.Println("Invalid appData format.")
		}
	}
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

type App struct {
	Script string `json:"script"`
}
