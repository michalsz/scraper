package main

import
(
	"fmt"
	"./page"
	"os"
	//"strconv"
)

func main(){
	url := os.Args[1]
	fmt.Println(url)
	wwwPage := page.Page{url, ""}
	for i := 0; i < 5; i++ {
	  wwwPage.ReadPage()
	  fmt.Println(wwwPage.PageBody())
    }
}
