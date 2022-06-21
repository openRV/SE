package desktop

// 设置文档可见性
// 只应用于docsId
// "POST"

type SetVisibilityParams struct {
	DocsId     string `json:"docsId"`
	Visibiliry string `json:"visibility"` // public | private
}

type SetVisibilityResult struct {
	Success bool `json:"success"`
}

type SetVisibility struct {
	Params SetVisibilityParams
	Result SetVisibilityResult
}
