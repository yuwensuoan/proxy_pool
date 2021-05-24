package models

// 代理模型
type ProxyModel struct {
	BaseModel
	Protocol string
	Proxy string
	Anonymous int
	CheckNum int
	FailCount int
	LastStatus int
	LastTime int
	Region string
	Source string
}

// 指定表名
func (ProxyModel) TableName() string {
	return "proxy"
}