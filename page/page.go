package page

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Page struct {
	Url string "xing.de"
}

func PageBody(page Page){
	resp, err := http.Get(page.Url)
	if err != nil{
		fmt.Println("error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("error")
	}
	fmt.Println(string(body))
}