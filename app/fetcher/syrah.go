package fetcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
	"time"
)

// 爬取西拉代理

type SyrahFetcher struct {
	BaseFetcher
}


func (S SyrahFetcher) Fetch(maxPage int) []map[string]interface{} {

	var data []map[string]interface{}

	ch := make(chan map[string]interface{}, 30)

	urlList := []string{
		"http://www.xiladaili.com/putong/%d/",
		"http://www.xiladaili.com/gaoni/%d/",
		"http://www.xiladaili.com/http/%d/",
		"http://www.xiladaili.com/https/%d/",
	}

	for _, baseUrl := range urlList {
		for i := 1; i <= maxPage; i++ {
			url := fmt.Sprintf(baseUrl, i)
			// fmt.Println(url)
			doc := S.fetchDocument(url, "utf8")
			go S.fetchData(doc, ch)
		}
		break
	}

	for	{
		select {
		case m := <-ch:
			data = append(data, m)
		case <-time.After(10 * time.Second):
			return data
		}
	}
}

func (S SyrahFetcher) fetchData (doc *goquery.Document, ch chan map[string]interface{}) []map[string]interface{} {

	var result []map[string]interface{}

	if doc == nil {
		 return result
	}
	reg, err := regexp.Compile("^[A-Za-z]+")
	if err != nil {
		return result
	}

	doc.Find("tbody>tr").Each(func(i int, selection *goquery.Selection){
		proxy := selection.Find("td:nth-child(1)").Text()
		protocol := reg.FindString(selection.Find("td:nth-child(2)").Text())
		// anonymity := selection.Find("td:nth-child(3)").Text() // 匿名度
		region := strings.Replace(selection.Find("td:nth-child(4)").Text(), " ", "/", 3)
		// timeout := selection.Find("td:nth-child(5)").Text()
		ch <- map[string]interface{}{
			"proxy": proxy,
			"protocol": strings.ToLower(protocol),
			"region": region,
			"source": "西拉代理",
		}
	})

	return result
}