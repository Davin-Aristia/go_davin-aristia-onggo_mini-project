package usecase

import (
	// "errors"

	"go-mini-project/dto"
	// "go-mini-project/middleware"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type CategoryUsecase interface {
	Get(name string) ([]dto.CategoryResponse, error)
	GetById(id int) (dto.CategoryResponse, error)
	Create(payloads dto.CategoryRequest) (dto.CategoryResponse, error)
	Update(payloads dto.CategoryRequest, id int) (dto.CategoryResponse, error)
	Delete(id int) error
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepository: categoryRepo}
}

func (s *categoryUsecase) Get(name string) ([]dto.CategoryResponse, error) {
	categoryData, err := s.categoryRepository.Get(name)
	if err != nil {
		return []dto.CategoryResponse{}, err
	}

	var categoryResponse []dto.CategoryResponse
	for _, category := range categoryData {
		categoryResponse = append(categoryResponse, dto.ConvertToCategoryResponse(category))
	}

	return categoryResponse, nil
}

func (s *categoryUsecase) GetById(id int) (dto.CategoryResponse, error) {
	categoryData, err := s.categoryRepository.GetById(id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	categoryResponse := dto.ConvertToCategoryResponse(categoryData)

	return categoryResponse, nil
}

func (s *categoryUsecase) Create(payloads dto.CategoryRequest) (dto.CategoryResponse, error) {
	categoryData := model.Category{
		Name : payloads.Name,
	}

	categoryData, err := s.categoryRepository.Create(categoryData)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	categoryResponse := dto.ConvertToCategoryResponse(categoryData)

	return categoryResponse, nil
}

func (s *categoryUsecase) Update(payloads dto.CategoryRequest, id int) (dto.CategoryResponse, error) {
	_, err := s.categoryRepository.GetById(id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	
	categoryData := model.Category{
		Name : payloads.Name,
	}

	err = s.categoryRepository.Update(categoryData, id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	categoryData, err = s.categoryRepository.GetById(id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	categoryResponse := dto.ConvertToCategoryResponse(categoryData)
	
	return categoryResponse, nil
}

func (s *categoryUsecase) Delete(id int) error {
	_, err := s.categoryRepository.GetById(id)
	if err != nil {
		return err
	}

	err = s.categoryRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
