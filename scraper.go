package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/michalsz/scraper/page"
	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	fmt.Println("Page to parse: ", url)
	wwwPage := page.Page{url, ""}
	wwwPage.ReadPage()
	// fmt.Println(wwwPage.Body)
	var chanNode chan page.Node
	var title string
	var a []string
	wantedTokens := []string{
		"a", "title",
	}

	chanNode = wwwPage.GetTokensFromBody(wantedTokens)

	for node := range chanNode {
		// fmt.Println(node.Token)
		if node.Type == "title" {
			fmt.Println(node.Token)
			tt := node.Doc.Next()

			if tt == html.TextToken {
				next := node.Doc.Token()
				title = strings.TrimSpace(next.Data)
			}
		}

		 if node.Type == "a" {
		 	fmt.Println(node.Token)
		 	a = append(a, node.Type)
		 }

	}

	fmt.Println("title", title)
	fmt.Println("a", a)
}
