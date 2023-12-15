package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/net/html"
)

func GetAttr(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

func IterateNodes(n *html.Node, f func(*html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		IterateNodes(c, f)
	}
}

func RandomUserAgents() string {

	file, err := os.Open("user-agent.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("File is empty")
	}

	r := rand.Intn(len(lines))

	return lines[r]
}

