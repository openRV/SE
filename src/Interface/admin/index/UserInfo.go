package index

// 用户运行状况
// /admin/userinfo

type UserInfoParams struct {
}

type D_UserIncreaseData struct {
	Date string
	Num  int
}

type M_UserIncreaseData struct {
	Month string
	Num   int
}

type UserInfoResult struct {
	Success bool         `json:"success"`
	Data    UserInfoData `json:"data"`
}

type UserInfoData struct {
	UserNumbers    int                  `json:"userNumbers"`
	OnlineNumbers  int                  `json:"onlineNumbers"`
	D_UserIncrease []D_UserIncreaseData `json:"onlineUsers"`
	M_UserIncrease []M_UserIncreaseData `json:"totalUsers"`
	// 按月新增
	// 按天新增
}

type UserInfo struct {
	Params UserInfoParams
	Result UserInfoResult
}
