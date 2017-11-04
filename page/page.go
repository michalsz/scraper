package page

import (
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
)


type Content interface{
	PageBody() string
	RedayPage() string
}

type Page struct {
	Url string ""
	Body string ""
}

func (page Page) PageBody() string{
	return page.Body
}

func (page *Page) ReadPage(){
	resp, err := http.Get(page.Url)
	if err != nil{
		fmt.Println("error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("error")
	}
	page.Body = string(body)
}

func IsHttps(page Page) bool{
	return strings.Contains(page.Url, "https")
}