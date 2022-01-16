package inMemoryTypes

type PostInMemoryRequest struct {
	Key   string `json:"key" `
	Value string `json:"value" `
}

type GetInMemoryResponse struct {
	Value string `json:"value"`
}
