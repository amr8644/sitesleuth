package main

import (
	"fmt"
	"time"
)

func main() {

	startTime := time.Now()
	Start()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Elapsed Time:", elapsedTime)
}
