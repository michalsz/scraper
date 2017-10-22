package main

import
(
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
  resp, err := http.Get("http://google.com/")
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
