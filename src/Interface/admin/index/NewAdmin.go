// @Title NewAdmin.go
// @Description 关于 NewAdmin 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

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
