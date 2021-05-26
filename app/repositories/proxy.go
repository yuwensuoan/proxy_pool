package repositories

import (
	"bytes"
	"fmt"
	"proxy_pool/app/models"
	"proxy_pool/database"
)

type ProxyRepository struct {
	BaseRepository
	Model models.ProxyModel
}

func (ProxyRepository) New() ProxyRepository {
	return ProxyRepository{}
}

// 批量插入数据
func (ProxyRepository) BatchInsert(models []map[string]interface{}) error {
	var buffer bytes.Buffer

	sql := "insert into `proxy` (`protocol`, `proxy`, `region`, `source`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}

	for i, model := range models {
		if i == len(models) - 1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s', '%s', '%s');", model["protocol"], model["proxy"], model["region"], model["source"]))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s'),", model["protocol"], model["proxy"], model["region"], model["source"]))
		}
	}
	return database.DB.Exec(buffer.String()).Error
}

// 获取proxy列表
func (P ProxyRepository) GetList(page int, size int) []map[string]interface{} {
	var proxies []map[string]interface{}
	database.DB.Model(&P.Model).Limit(size).
		Select("id,protocol,proxy,region,source").Offset((page - 1) * size).Find(&proxies)
	return proxies
}


// 获取一条记录
func (P ProxyRepository) GetFirst() map[string]interface{} {
	proxy := map[string]interface{}{}

	database.DB.Model(&P.Model).Select("id,protocol,proxy,region,source").Take(&proxy)

	return proxy
}

// 判断记录是否存在
func (P ProxyRepository) IsExists(proxy string) bool {
	obj := map[string]interface{}{}
	database.DB.Model(&P.Model).Where("proxy = ?", proxy).Take(&obj)
	if len(obj) < 1 {
		return false
	}
	return true
}

// 根据查找记录
func (P ProxyRepository) GetFirstById(proxy string) map[string]interface{} {
	obj := map[string]interface{}{}

	database.DB.Model(&P.Model).Where("proxy = ?", proxy).Take(&obj)

	return obj
}

// 创建记录
func (P ProxyRepository) Create(record map[string]interface{}) {
	database.DB.Create(&models.ProxyModel{
		Proxy: record["proxy"].(string),
		Protocol: record["protocol"].(string),
		Region:  record["region"].(string),
		Source: record["source"].(string),
	})
}