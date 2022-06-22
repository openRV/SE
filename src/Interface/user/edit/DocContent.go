package edit

// /user/readcontent
// "GET"

type DocContentParams struct {
	DocId string `json:"docsId"`
}

type DocData struct {
	DocContent string `json:"docsContent"`
}

type DocContentResult struct {
	Success bool    `json:"success"`
	Data    DocData `json:"data"`
}

type DocContent struct {
	Params DocContentParams
	Result DocContentResult
}
