package request

type ChatGPTRequest struct {
	Texto string `json:"text"`
}

type ChatGPTRequestRepository interface {
	Question(cgpt *ChatGPTRequest) (*ChatGPTRequest, error)
}
