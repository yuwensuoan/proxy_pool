package main

import (
	"fmt"
	"proxy_pool/app/fetcher"
	"proxy_pool/app/repositories"
)

func main()  {
	// boostrap.Server.Run(":8080")
	fetch := fetcher.CloudFetcher{}.NewFetcher(&repositories.ProxyRepository{})
	data := fetch.Fetch(1)
	fmt.Println(data)

}