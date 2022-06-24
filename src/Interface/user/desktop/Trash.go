// @Title Trash.go
// @Description 关于 Trash 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 回收站 30天清理 只含文件
// /user/trash
// "GET"

type TrashParams struct {
	CurPage  int `json:"curPage"`  // default 1
	PageSize int `json:"pageSize"` // default 15
}

type TrashData struct {
	DocsName   string `json:"docsName"`
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
