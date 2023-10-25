package dto

type ChatbotRequest struct {
	Content string `json:"content" form:"content"`
}