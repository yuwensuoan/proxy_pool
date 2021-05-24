package services

import (
	"proxy_pool/app/repositories"
)

// proxy相关业务逻辑
type ProxyService struct {
	BaseService
	Repository repositories.ProxyRepository
}

// 返回实例
func (ProxyService) New() ProxyService {
	return ProxyService{}
}

// 获取代理列表
func (P ProxyService) GetList() []map[string]interface{} {
	data := P.Repository.GetList(1, 10)
	return data
}

// 获取第一条代理
func (P ProxyService) GetFirst() map[string]interface{} {
	data := P.Repository.GetFirst()
	return data
}