package index

// Register
// /open/register
// "POST"

type RegisterParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterResult struct {
	Success string `json:"success"`
}

type Register struct {
	Params RegisterParams
	Result RegisterResult
}
