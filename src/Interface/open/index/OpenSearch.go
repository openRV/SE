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

type OpenSearchResult struct {
	Success bool `json:"success"`
	Total   int  `json:"total"` // 总共的数量
	Data    []OpenSearchData `json:"data"`
}

type OpenSearchData struct {
		DocsId     string `json:"docsId"` // docs key
		DocsName   string `json:"docsName"`
		Author     string `json:"author"`
		LastUpdate string `json:"lastUpdate"` // yyyy-mm-dd
		ViewCounts int    `json:"viewCounts"`

}

type OpenSearch struct {
	Params OpenSearchParams
	Result OpenSearchResult
}
