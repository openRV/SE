package desktop

// 获取文件夹内容
// /user/dircontent
// "GET"

type DirContentParams struct {
	DirId    string `json:"DirId"`
	CurPage  int    `json:"curPage"`
	PageSize int    `json:"pageSize"` // 文件夹与文件一样占一行
}

type DirContentData struct {
	Dir  []DirListItem  `json:"dir"`
	Docs []DocsListItem `json:"docs"`
}

type DirContentResult struct {
	Success bool           `json:"success"`
	Total   int            `json:"total"`
	Data    DirContentData `json:"data"`
}

type DirContent struct {
	Params DirContentParams
	Result DirContentResult
}
