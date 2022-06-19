package desktop

// 获取文件夹内容
// /user/dircontent
// "GET"

type DirContentParams struct {
	DocsId   string `json:"docsId"`
	CurPage  int    `json:"curPage"`
	PageSize int    `json:"pageSize"` // 文件夹与文件一样占一行
}

type DirContentResult struct {
	Success bool `json:"success"`
	Total   int  `json:"total"`
	Data    []struct {
		Dir  DirListItem  `json:"dir"`
		Docs DocsListItem `json:"docs"`
	} `json:"data"`
}

type DirContent struct {
	Params DirContentParams
	Result DirContentResult
}
