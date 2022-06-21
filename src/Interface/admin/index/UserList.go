package index

// /admin/userlist
// "GET"
type UserListParams struct {
	CurPage  int `json:"curPage"`
	PageSize int `json:"pageSize"`
}

type UserListResult struct {
	Success bool        `json:"success"`
	Total   int         `json:"total"`
	Data    [][3]string `json:"data"`
}

type UserList struct {
	Params UserListParams
	Result UserListResult
}
