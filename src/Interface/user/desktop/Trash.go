package desktop

// 回收站 30天清理 只含文件
// /user/trash
// "GET"

type TrashParams struct {
	CurPage  int `json:"curPage"`  // default 1
	PageSize int `json:"pageSize"` // default 15
}

type TrashData struct {
	DocsId     string `json:"docsId"`
	Author     string `json:"author"`
	DeleteDate string `json:"deleteDate"`
}

type TrashResult struct {
	Success bool        `json:"success"`
	Total   int         `json:"total"`
	Data    []TrashData `json:"data"`
}

type Trash struct {
	Params TrashParams
	Result TrashResult
}
