package index

// 系统运行状况
// /admin/storageinfo
// "GET"

type StorageInfoParams struct {
}

type StorageInfoResult struct {
	Success bool `json:"success"`
	Data    struct {
		TotalStorage string `json:"totalStorage"` // 带单位
		UsingStorage string `json:"usingStorage"`
		StoreRate    string `json:"storeRate"` // 使用率
		D_increase   []D_increaseData
		M_increase   []M_increaseData
		// 此处我不清楚画图需要的对应的数据类型
		// 励巨画图时弄一下?
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
