package main

import
(
	"fmt"
	"./page"
)

func main(){
  p := page.Page{"http://xing.com"}
  fmt.Println(p.Url)
  page.PageBody(p)
}
