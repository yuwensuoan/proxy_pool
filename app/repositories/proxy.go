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
		Select("id,protocol,proxy,region,source,check_num,fail_count,last_status,last_time").Offset((page - 1) * size).Find(&proxies)
	return proxies
}


// 获取一条记录
func (P ProxyRepository) GetFirst() map[string]interface{} {
	proxy := map[string]interface{}{}

	database.DB.Model(&P.Model).Take(&proxy)

	return proxy
}