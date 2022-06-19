package desktop

// 最近查看记录 只需返回文件类型
// /user/lastview
// "GET"

type LastViewParams struct {
	CurPage  int `json:"curPage"`
	PageSize int `json:"pageSize"`
}

type LastViewResult struct {
	Success bool           `json:"success"`
	Total   int            `json:"total"`
	Data    []DocsListItem `json:"data"`
}

type LastView struct {
	Params LastViewParams
	Result LastViewResult
}
