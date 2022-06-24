// @Title NewDir.go
// @Description 关于 NewDir 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

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
