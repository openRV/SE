// @Title UserSearch.go
// @Description 关于 UserSearch 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 查找当前用户可见的文档
// /user/search
// "GET"

type UserSearchParams struct {
	SearchContent string `json:"searchContent"` // 搜索内容
	SearchType    string `json:"searchType"`    // Title | Author
	CurPage       int    `json:"curPage"`       // default 1
	PageSize      int    `json:"pageSize"`      // default 15
}

type UserSearchResult struct {
	Success bool           `json:"success"`
	Total   int            `json:"total"`
	Data    []DocsListItem `json:"data"`
}

type UserSearch struct {
	Params UserSearchParams
	Result UserSearchResult
}
