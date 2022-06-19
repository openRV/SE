package index

// 设置自己的信息
// /admin/setinfo
// "POST"

type SetInfoParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type SetInfoResult struct {
	Success bool `json:"success"`
}
