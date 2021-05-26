package models

// 代理模型
type ProxyModel struct {
	BaseModel
	Protocol string
	Proxy string `gorm:"not null;index;comment:代理IP和端口"`
	Region string
	Source string
}

// 指定表名
func (ProxyModel) TableName() string {
	return "proxy"
}