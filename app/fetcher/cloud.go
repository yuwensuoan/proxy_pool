package fetcher

/**
获取云代理ip列表
 */
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

type CloudFetcher struct {
	BaseUrl string
	BaseFetcher
}

// 获取指定页数的代理
func (F CloudFetcher) Fetch(maxPage int) []map[string]interface{} {
	F.BaseUrl = "http://www.ip3366.net/free/?stype=%d&page=%d"

	var result []map[string]interface{}

	ch := make(chan map[string]interface{}, 20)

	proxyTypeList := []int{1, 2} // 1: 高匿, 2: 透明

	for _, proxyType := range proxyTypeList {
		//url := fmt.Sprintf(F.BaseUrl, proxyType, 1)
		//doc := F.fetchDocument(url)
		//max := F.getMaxPage(doc)
		//if max > maxPage {
		//	maxPage = max
		//}
		for i:=1; i <= maxPage; i++ {
			url := fmt.Sprintf(F.BaseUrl, proxyType, i)
			doc := F.fetchDocument(url, "gbk")
			go F.fetchData(doc, ch)
		}
	}

	for	{
		select {
		case m := <-ch:
			result = append(result, m)
		case <-time.After(10 * time.Second):
			return result
		}
	}
}


// 获取详细代理数据
func (F CloudFetcher) fetchData (doc *goquery.Document, ch chan map[string]interface{}) []map[string]interface{} {

	var result []map[string]interface{}

	if doc == nil {
		return result
	}

	// 查找代理列表
	doc.Find("tbody>tr").Each(func(i int, selection *goquery.Selection) {
		host := selection.Find("td:nth-child(1)").Text()  // 主机
		port := selection.Find("td:nth-child(2)").Text()  // 端口
		//anonymity := selection.Find("td:nth-child(3)").Text() // 匿名度
		protocol := selection.Find("td:nth-child(4)").Text() // 协议
		//method := selection.Find("td:nth-child(5)").Text()  // 请求方法
		region := selection.Find("td:nth-child(5)").Text() // 地区
		//timeout := selection.Find("td:nth-child(7)").Text() // 响应时间
		//lastTime := selection.Find("td:nth-child(8)").Text() // 最后检测时间
		region = strings.Replace(region, "SSL_", "", 1)
		ch<- map[string]interface{}{
			"proxy": host + ":" + port,
			"protocol": strings.ToLower(protocol),
			"region": region,
			"source": "云代理",
		}
	})
	return result
}


//// 获取最大页数
//func (F CloudFetcher) getMaxPage(doc *goquery.Document) int {
//	href, exists := doc.Find("#listnav > ul > a:nth-child(10)").Attr("href")
//	if !exists {
//		return 1
//	}
//	strList := strings.Split(href, "=")
//	if len(strList) <= 0 {
//		return 1
//	}
//
//	maxPage, err:= strconv.Atoi(strList[len(strList)-1])
//	if err != nil {
//		return 1
//	}
//
//	return maxPage
//}

