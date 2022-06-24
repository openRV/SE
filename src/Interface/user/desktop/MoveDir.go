// @Title MoveDir.go
// @Description 关于 MoveDir 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 移动
// /user/movedir
// "POST"

type MoveDirParams struct {
	DirId  string `json:"dirId"`  //
	MoveTo string `json:"moveTo"` // a dir
}

type MoveDirResult struct {
	Success bool `json:"success"`
}

type MoveDir struct {
	Params MoveFileParams
	Result MoveFileResult
}
