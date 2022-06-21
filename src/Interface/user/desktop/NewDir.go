package desktop

// 新建文件夹
// "POST"

type NewDirParams struct {
	FatherDirId string `json:"fatherDirId"` // 默认根文件夹
	DirName     string `json:"dirName"`
}

type NewDirResult struct {
	Success bool `json:"success"`
}

type NewDir struct {
	Params NewDirParams
	Result NewDirResult
}
