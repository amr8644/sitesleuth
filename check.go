package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)


var result map[string]interface{}

func ConvertGolang(key string) []string {

	var n_p []string
	pattern := `"` + key + `"\s*:\s*"(.*?)"`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(ReadFile(), -1)

	if len(matches) == 0 {
		fmt.Println("Key not found in JSON file.")
	}

	var values []string
	for _, match := range matches {
		if len(match) >= 2 {
			values = append(values, match[1]) // match[1] contains the captured value
		}
	}

	for _, v := range values {
		escapedValue := regexp.QuoteMeta(v)
		htmlPattern := `(?s)` + escapedValue
		fmt.Println(htmlPattern)
		n_p = append(n_p, htmlPattern)
	}

	return n_p
}

func CheckHTML(html string) {


	var ans []string
	re := ConvertGolang("script")

	for _, v := range re {
        fmt.Println(v)
		r := regexp.MustCompile(v)
		ans = r.FindStringSubmatch(html)
		fmt.Println(ans)
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
