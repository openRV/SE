// @Title WhiteDocs.go
// @Description 关于 WhiteDocs 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package edit

// /user/writecontent
// "POST"

type WriteDocsParams struct {
	DocsId      string `json:"docsId"`
	DocsContent string `json:"docsContent"`
}

type WriteDocsResult struct {
	Success bool `json:"success"`
}

type WriteDocs struct {
	Params WriteDocsParams
	Result WriteDocsResult
}
