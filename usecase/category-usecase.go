package usecase

import (
	// "errors"

	"go-mini-project/dto"
	// "go-mini-project/middleware"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type CategoryUsecase interface {
	Get(name string) ([]model.Category, error)
	GetById(id int) (model.Category, error)
	Create(payloads dto.CategoryRequest) (model.Category, error)
	Update(payloads dto.CategoryRequest, id int) (model.Category, error)
	Delete(id int) error
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepository: categoryRepo}
}

func (s *categoryUsecase) Get(name string) ([]model.Category, error) {
	categoryData, err := s.categoryRepository.Get(name)
	if err != nil {
		return []model.Category{}, err
	}

	return categoryData, nil
}

func (s *categoryUsecase) GetById(id int) (model.Category, error) {
	categoryData, err := s.categoryRepository.GetById(id)
	if err != nil {
		return model.Category{}, err
	}

	return categoryData, nil
}

func (s *categoryUsecase) Create(payloads dto.CategoryRequest) (model.Category, error) {
	categoryData := model.Category{
		Name : payloads.Name,
	}

	categoryData, err := s.categoryRepository.Create(categoryData)
	if err != nil {
		return model.Category{}, err
	}
	return categoryData, nil
}

func (s *categoryUsecase) Update(payloads dto.CategoryRequest, id int) (model.Category, error) {
	_, err := s.categoryRepository.GetById(id)
	if err != nil {
		return model.Category{}, err
	}
	
	categoryData := model.Category{
		Name : payloads.Name,
	}

	err = s.categoryRepository.Update(categoryData, id)
	if err != nil {
		return model.Category{}, err
	}

	categoryData, err = s.categoryRepository.GetById(id)
	if err != nil {
		return model.Category{}, err
	}
	
	return categoryData, nil
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
