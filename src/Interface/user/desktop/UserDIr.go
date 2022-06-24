// @Title UserDir.go
// @Description 关于 UserDir 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 获取文件树
// /user/dir
// "GET"

type UserDirParams struct {
	// UserName string `json:"userName"` // 有token信息是不是不需要这个
}

type UserDirResult struct {
	Success bool  `json:"success"`
	Data    []Dir `json:"data"`
}

type UserDir struct {
	Params UserDirParams
	Result UserDirResult
}
