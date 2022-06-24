// @Title SetVisibility.go
// @Description 关于 SetVisibility 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

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
