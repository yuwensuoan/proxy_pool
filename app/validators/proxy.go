package validators

type ProxyGetList struct {
	Page int `form:"page" json:"page" binding:"omitempty"`
	PageSize int `form:"page_size" json:"page_size" binding:"omitempty"`
}

type ProxyGetFirst struct {

}