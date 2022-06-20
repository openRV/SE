package index

// /open/login
// "POST"

type LoginParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResult struct {
	Success bool `json:"success"`
	Data    struct {
		Role   string `json:"role"` // admin | user
		Token  string `json:""token`
		Name   string `json:"name"`   // 用户 Account 唯一
		Avatar string `json:"avatar"` // 头像 string
	} `json:"data"`
}

type Login struct {
	Params LoginParams
	Result LoginResult
}
