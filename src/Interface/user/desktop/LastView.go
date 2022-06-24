// @Title LastView.go
// @Description 关于 LastView 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

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
