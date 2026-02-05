package crawlers

import "fmt"
import "github.com/Masud2017/proxy_crawler/models"

type FreeProxyListCrawler struct {}

func (freeProxyListCrawler FreeProxyListCrawler) Crawl() {
	var data models.CrawledData = models.CrawledData{}
	
	fmt.Println("Hello world this is my test data: ",data)
}

