package usecase

import (
	"context"
	"fmt"
	"strconv"

	"go-mini-project/dto"
	"go-mini-project/repository"

	openai "github.com/sashabaranov/go-openai"
)

type ChatbotUsecase interface {
    GetLaptopRecommendation(request dto.ChatbotRequest) (string, error)
}

type chatbotUsecase struct {
    chatbotRepository repository.ChatbotRepository
    categoryRepository repository.CategoryRepository
    bookRepository repository.BookRepository
}

func NewChatbotUsecase(chatbotRepo repository.ChatbotRepository, categoryRepo repository.CategoryRepository, bookRepo repository.BookRepository) *chatbotUsecase {
    return &chatbotUsecase{
		chatbotRepository: chatbotRepo,
		categoryRepository: categoryRepo,
		bookRepository: bookRepo,
	}
}

func (s *chatbotUsecase) GetLaptopRecommendation(request dto.ChatbotRequest) (string, error) {
	categoryData, err := s.categoryRepository.Get("")
	if err != nil {
		return "", err
	}

    // content := "Give my customer book recommendation from my store. Here is the list of books in my store with its categories:"
    content := "I am a customer from certain book store that have list of books:"

    for _, category := range categoryData {
        content += "\ncategory " + category.Name + ":"
        for index, book := range category.Books {
            content += "\n" + strconv.Itoa(index) + ". Book: " + book.Title + ", Author: " + book.Author + ", Price: Rp. " + strconv.Itoa(int(book.Price))
        }
    }

    // content += "\n\nand here is my customer chat: " + request.Content
    content += "\n" + request.Content

    fmt.Println("content: ", content)

    ctx := context.Background()
    messages := []openai.ChatCompletionMessage{
        {
            Role:    openai.ChatMessageRoleSystem,
            Content: "You are a book recommendation chatbot.",
        },
        {
            Role:    openai.ChatMessageRoleUser,
            Content: content,
        },
    }
    model := openai.GPT3Dot5Turbo
    resp, err := s.chatbotRepository.GetCompletionFromMessages(ctx, messages, model)
    result := resp.Choices[0].Message.Content
    return result, err
}
