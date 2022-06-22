package desktop

// 删除 相当于移动到回收站
// /user/deleleitem
// "DELETE"

type DeleteItemParams struct {
	DocsId string `json:"docsId"`
	IsDir  string `json:"isDir"`
}

type DeleteItemResult struct {
	Success bool `json:"success"`
}

type DeleteItem struct {
	Params DeleteItemParams
	Result DeleteItemResult
}
