package fetcher

/**
获取云代理ip列表
 */
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"net/http"
	"strconv"
	"strings"
)

type CloudFetcher struct {
	BaseUrl string
}

// 获取指定页数的代理
func (F CloudFetcher) Fetch(totalPage int) []map[string]interface{} {
	F.BaseUrl = "http://www.ip3366.net/free/?stype=%d&page=%d"

	var result []map[string]interface{}

	proxyTypeList := []int{1, 2} // 1: 高匿, 2: 透明

	for _, proxyType := range proxyTypeList {
		url := fmt.Sprintf(F.BaseUrl, proxyType, 1)
		doc := F.fetchDocument(url)
		maxPage := F.getMaxPage(doc)
		if maxPage > totalPage {
			maxPage = totalPage
		}
		for i:=1; i <= maxPage; i++ {
			url := fmt.Sprintf(F.BaseUrl, proxyType, i)
			doc := F.fetchDocument(url)
			data := F.fetchData(doc)
			result = append(result, data...)
		}
	}
	return result
}


// 获取文档
func (F CloudFetcher) fetchDocument(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}

	reader := mahonia.NewDecoder("gbk").NewReader(resp.Body)

	doc, err := goquery.NewDocumentFromReader(reader)

	return doc
}

// 获取最大页数
func (F CloudFetcher) getMaxPage(doc *goquery.Document) int {
	href, exists := doc.Find("#listnav > ul > a:nth-child(10)").Attr("href")
	if !exists {
		return 1
	}
	strList := strings.Split(href, "=")
	if len(strList) <= 0 {
		return 1
	}

	maxPage, err:= strconv.Atoi(strList[len(strList)-1])
	if err != nil {
		return 1
	}

	return maxPage
}

// 获取详细代理数据
func (F CloudFetcher) fetchData (doc *goquery.Document) []map[string]interface{} {

	var result []map[string]interface{}

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

		result = append(result, map[string]interface{}{
			"proxy": host + ":" + port,
			"protocol": strings.ToLower(protocol),
			"region": region,
			"source": "cloud",
		})
	})
	return result
}