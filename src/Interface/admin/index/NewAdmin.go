package index

// 新建管理员账户
// /admin/newadmin
// "POST"

type NewAdminParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type NewAdminResult struct {
	Success bool `json:"success"`
}
