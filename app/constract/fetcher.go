package constract

import "proxy_pool/app/models"

type Fetcher interface {
	Fetch(totalPage int) []models.ProxyModel
}