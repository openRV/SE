// @Title MoveFile.go
// @Description 关于 MoveFile 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 移动
// /user/movefile
// "POST"

type MoveFileParams struct {
	DocsId string `json:"docsId"` //
	MoveTo string `json:"moveTo"` // a dir
}

type MoveFileResult struct {
	Success bool `json:"success"`
}

type MoveFile struct {
	Params MoveFileParams
	Result MoveFileResult
}
