package main

import (
	"fmt"
	"time"
)

func main() {
	// Record the start time
	startTime := time.Now()
	Start()

	// Record the end time
	endTime := time.Now()

	// Calculate the elapsed time
	elapsedTime := endTime.Sub(startTime)

	// Print the elapsed time
	fmt.Println("Elapsed Time:", elapsedTime)
	//PrintData()
}
