package repository

import (
    "context"
    openai "github.com/sashabaranov/go-openai"
)

type ChatbotRepository interface {
    GetCompletionFromMessages(ctx context.Context, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error)
}

type chatbotRepository struct {
    client *openai.Client
}

func NewChatbotRepository(client *openai.Client) *chatbotRepository {
    return &chatbotRepository{client}
}

func (r *chatbotRepository) GetCompletionFromMessages(ctx context.Context, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
    if model == "" {
        model = openai.GPT3Dot5Turbo
    }

    resp, err := r.client.CreateChatCompletion(
        ctx,
        openai.ChatCompletionRequest{
            Model:    model,
            Messages: messages,
        },
    )
    return resp, err
}
