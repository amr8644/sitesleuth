package main

import (
	"fmt"
	"os"
)


func CheckHTML(html string){

// Open the JSON file
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

    fmt.Println(file)


}
