package edit

// /user/writecontent
// "POST"

type WriteDocsParams struct {
	DocsId      string `json:"docsId"`
	DocsContent string `json:"docsContent"`
}

type WriteDocsResult struct {
	Success bool `json:"success"`
}

type WriteDocs struct {
	Params WriteDocsParams
	Result WriteDocsResult
}
