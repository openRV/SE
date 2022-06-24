// @Title StorageInfo.go
// @Description 关于 StorageInfo 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package index

// 系统运行状况
// /admin/storageinfo
// "GET"

type StorageInfoParams struct {
}

type StorageInfoResult struct {
	Success bool `json:"success"`
	Data    struct {
		TotalStorage float32 `json:"totalStorage"`
		UsingStorage float32 `json:"usingStorage"`
		D_increase   []D_increaseData
		M_increase   []M_increaseData
	} `json:"data"`
}

type D_increaseData struct {
	Date string
	Num  int
}

type M_increaseData struct {
	Month string
	Num   int
}
