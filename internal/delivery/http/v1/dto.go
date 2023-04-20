package v1

type CreateShorteningRequest struct {
	URL string `json:"url,omitempty"`
}

type CreateShorteningResponse struct {
	Short string `json:"short,omitempty"`
}

type RestoreResponse struct {
	Original string `json:"original,omitempty"`
}

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
