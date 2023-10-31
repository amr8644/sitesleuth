package main

import "golang.org/x/net/html"



func GetAttr(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

// Iterate DOM nodes
func IterateNodes(n *html.Node, f func(*html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		IterateNodes(c, f)
	}
}
