package index

// /open/hotdocs
// "GET"

type HotdocsParams struct {
}

type HotdocsResult struct {
	Success bool          `json:"success"`
	Total   int           `json:"total"`
	Data    []HotdocsData `json:"data"`
}

type HotdocsData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	DownloadNum int    `json:"downloadNum"`
}

type Hotdocs struct {
	Params HotdocsParams
	Result HotdocsResult
}
