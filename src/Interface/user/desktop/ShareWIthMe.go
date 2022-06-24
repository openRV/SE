// @Title ShareWithMe .go
// @Description 关于 ShareWithMe 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 与我共享 只含文件
// /user/sharewithme
// "GET"

type ShareWithMeParams struct {
	CurPage  int `json:"curPage"`  // default 1
	PageSize int `json:"pageSize"` // default 15
}

type ShareWithMeResult struct {
	Success bool           `json:"success"`
	Total   int            `json:"total"`
	Data    []DocsListItem `json:"data"`
}

type ShareWithMe struct {
	Params ShareWithMeParams
	Result ShareWithMeResult
}
