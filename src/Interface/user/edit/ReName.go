// @Title ReName.go
// @Description 关于 ReName 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package edit

// 更改名称
// /user/rename
// "POST"

type ReNameParams struct {
	DocsId  string `json:"docsId"`
	NewName string `json:"newName"`
	IsDir   bool   `json:"isDir"`
}

type ReNameResult struct {
	Success bool `json:"success"`
}

type ReName struct {
	Params ReNameParams
	Result ReNameResult
}
