package services

import (
	"proxy_pool/app/fetcher"
	"proxy_pool/app/repositories"
)

type ProxyService struct {
	BaseService
}

// 获取代理列表
func (ProxyService) GetList() {
	fetch := fetcher.NewFetcher(&repositories.ProxyRepository{})
	fetch.FetchCloudProxy()
}

// 获取第一条代理
func (ProxyService) GetFirst() {

}