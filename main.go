package main

import (
)


func main() {
    response :=SendRequests("https://shielder.it")
    ScrapeURL(*response, "https://shielder.it")

}
