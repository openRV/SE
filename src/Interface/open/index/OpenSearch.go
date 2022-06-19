package index

// search for open docs
// 考虑名称与作者名
// /open/search
// "GET"

type OpenSearchParams struct {
	SearchContent string `json:"searchContent"`
	SearchType    string `json:"searchType"` // Title | Author
	CurPage       int    `json:"curPage"`    // default 1
	PageSize      int    `json:"pageSize"`   // default 15
}

type OpenSearchResult[T int | string] struct {
	Success bool `json:"success"`
	Total   int  `json:"total"` // 总共的数量
	Data    []struct {
		DocsId     string `json:"docsId"` // docs key
		DocsName   string `json:"docsName"`
		Author     string `json:"author"`
		LastUpdata string `json:"lastUpdate"` // yyyy-mm-dd
		ViewCounts T      `json:"viewCounts"` // TODO
	} `json:"data"`
}

type OpenSearch struct {
	Params OpenSearchParams
	Result OpenSearchResult[int] // TODO
}
