package main



func main() {
    tar := "https://www.w3schools.com/"
    response :=SendRequests(tar)
    ScrapeURL(*response,tar )

}
