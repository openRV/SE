// @Title DeleteUser.go
// @Description 关于 DeleteUser 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package index

// 删除用户 这里只看到普通用户
// /admin/deleteuser
// "DELETE"

type DeleteUserParams struct {
	UserName string `json:"userName"`
}

type DeleteUserResult struct {
	Success bool `json:"success"`
}

type DeleteUser struct {
	Params DeleteUserParams
	Result DeleteUserResult
}
