// @Title Interface.go
// @Description 关于 文件、文件夹、文件树的数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 定义文档列表项类型
type DocsListItem struct {
	DocsId     string `json:"docsId"`
	DocsName   string `json:"docsName"`
	Author     string `json:"author"`
	CreateDate string `json:"createDate"` // date type
	LastView   string `json:"lastView"`   // date type
	DocsType   string `json:"docsType"`   // private | public | cooperate
}

// 定义文件夹类型 单项
type DirListItem struct {
	DirId      string `json:"dirId"`
	DirName    string `json:"dirName"`
	Owner      string `json:"owner"`
	LastView   string `json:"lastView"`   // date
	CreateDate string `json:"createDate"` // date
}

// 文件树
type Dir struct {
	DirId   string `json:"dirId"`
	DirName string `json:"dirName"`
	Subdir  []Dir  `json:"subdir"`
}
