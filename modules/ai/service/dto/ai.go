package dto

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AiMsg struct {
	Messages  []Message `json:"messages"`
	Platform  string    `json:"platform"`
	ModelName string    `json:"modelName"`
}
