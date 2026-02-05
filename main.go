package main

import "fmt"
import "os"
import "reflect"
//import "github.com/anaskhan96/soup"
import "github.com/Masud2017/proxy_crawler/crawlers"
//import "github.com/Masud2017/proxy_crawler/utils"

func main() {
	resp, err := soup.Get("https://free-proxy-list.net/en/")
	if (err != nil) {
		os.Exit(1)
	}
	
	fmt.Println("Hello world checking the type of args: ",reflect.TypeOf(os.Args))

	crawlers.Crawl()
//	fmt.Println(resp)




}


