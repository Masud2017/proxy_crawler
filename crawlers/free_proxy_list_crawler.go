package crawlers

import "fmt"
import "github.com/Masud2017/proxy_crawler/models"
import "github.com/anaskhan96/soup"

type FreeProxyListCrawler struct {}

func (freeProxyListCrawler FreeProxyListCrawler) Crawl() {
	var data models.CrawledData = models.CrawledData{}
	var url := ""
	resp, err := soup.Get("https://free-proxy-list.net/en/")
        if (err != nil) {
		fmt.Println("There was a problem during parsing the html content. Please check the free-proxy-list crawler.")
                os.Exit(1)
        }
	
	fmt.Println("Hello world this is my test data: ",data)
}

