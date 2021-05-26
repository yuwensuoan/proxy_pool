package services

import (
	"proxy_pool/app/repositories"
	"proxy_pool/app/utils"
	"strconv"
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
func (P ProxyService) GetList(page string, pageSize string) utils.Response {
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		ps = 10
	}
	data := P.Repository.GetList(p, ps)
	return utils.Response{}.Success(data)
}

// 获取第一条代理
func (P ProxyService) GetFirst() utils.Response {
	data := P.Repository.GetFirst()
	return utils.Response{}.Success(data)
}