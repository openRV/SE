package index

// /admin/userlist
// "GET"
type UserListParams struct {
	CurPage  int `json:"curPage"`
	PageSize int `json:"pageSize"`
}

type UserListResult struct {
	Success bool `json:"success"`
	Total   int  `json:"total"`
	Data    struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		// 用户的下载记录与最近查看记录通过user接口访问
		Storage string `json:"storage"` //  占用存储空间大小
	} `json:"data"`
}

type UserList struct {
	Params UserListParams
	Result UserListResult
}
