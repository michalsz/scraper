package page

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Content interface {
	PageBody() string
	RedayPage() string
}

type Page struct {
	Url  string ""
	Body string ""
}

type Node struct {
	Type  string
	Token html.Token
	Doc   *html.Tokenizer
}

func (page *Page) ReadPage() {
	resp, err := http.Get(page.Url)
	if err != nil {
		fmt.Println("error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error")
	}
	page.Body = string(body)
}

func (page *Page) GetTokensFromBody(wantedTokens []string) (c chan Node) {
	c = make(chan Node)

	go func() {
		defer close(c)

		// https://play.golang.org/p/0MRSefJ_-E
		r := strings.NewReader(page.Body)
		z := html.NewTokenizer(r)

		for {
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				// End of the document, we're done
				return
			case tt == html.StartTagToken:
				token := z.Token()

				for _, name := range wantedTokens {
					if token.Data == name {
						c <- Node{token.Data, token, z}
					}
					continue
				}
			}
		}
	}()

	return c
}

func IsHttps(page Page) bool {
	return strings.Contains(page.Url, "https")
}

func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}
