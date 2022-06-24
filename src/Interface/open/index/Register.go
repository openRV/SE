// @Title Register.go
// @Description 关于 Register 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package index

// Register
// /open/register
// "POST"

type RegisterParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterResult struct {
	Success bool `json:"success"`
}

type Register struct {
	Params RegisterParams
	Result RegisterResult
}
