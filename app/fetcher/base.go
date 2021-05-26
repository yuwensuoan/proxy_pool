package fetcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"math/rand"
	"net/http"
	"time"
)

type BaseFetcher struct {

}

func (BaseFetcher) getUserAgent() string {
	userAgentList := []string{
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
	}
	return userAgentList[rand.Intn(len(userAgentList))]
}

// 获取文档
func (B BaseFetcher) fetchDocument(url string, unicode string) *goquery.Document {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", B.getUserAgent())
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	client := &http.Client{
		Timeout:   time.Second * 20, //超时时间
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	reader := mahonia.NewDecoder(unicode).NewReader(resp.Body)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return doc
}