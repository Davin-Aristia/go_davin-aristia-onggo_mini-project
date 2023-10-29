package usecase

import (

	"go-mini-project/dto"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type VendorUsecase interface {
	Get(name string) ([]dto.VendorResponse, error)
	GetById(id int) (dto.VendorResponse, error)
	Create(payloads dto.VendorRequest) (dto.VendorResponse, error)
	Update(payloads dto.VendorRequest, id int) (dto.VendorResponse, error)
	Delete(id int) error
}

type vendorUsecase struct {
	vendorRepository repository.VendorRepository
}

func NewVendorUsecase(vendorRepo repository.VendorRepository) *vendorUsecase {
	return &vendorUsecase{vendorRepository: vendorRepo}
}

func (s *vendorUsecase) Get(name string) ([]dto.VendorResponse, error) {
	vendorData, err := s.vendorRepository.Get(name)
	if err != nil {
		return []dto.VendorResponse{}, err
	}

	var vendorResponse []dto.VendorResponse
	for _, vendor := range vendorData {
		vendorResponse = append(vendorResponse, dto.ConvertToVendorResponse(vendor))
	}

	return vendorResponse, nil
}

func (s *vendorUsecase) GetById(id int) (dto.VendorResponse, error) {
	vendorData, err := s.vendorRepository.GetById(id)
	if err != nil {
		return dto.VendorResponse{}, err
	}

	vendorResponse := dto.ConvertToVendorResponse(vendorData)

	return vendorResponse, nil
}

func (s *vendorUsecase) Create(payloads dto.VendorRequest) (dto.VendorResponse, error) {
	vendorData := model.Vendor{
		Name : payloads.Name,
		Address : payloads.Address,
		Email : payloads.Email,
		PhoneNumber : payloads.PhoneNumber,
	}

	vendorData, err := s.vendorRepository.Create(vendorData)
	if err != nil {
		return dto.VendorResponse{}, err
	}

	vendorResponse := dto.ConvertToVendorResponse(vendorData)

	return vendorResponse, nil
}

func (s *vendorUsecase) Update(payloads dto.VendorRequest, id int) (dto.VendorResponse, error) {
	_, err := s.vendorRepository.GetById(id)
	if err != nil {
		return dto.VendorResponse{}, err
	}
	
	vendorData := model.Vendor{
		Name : payloads.Name,
		Address : payloads.Address,
		Email : payloads.Email,
		PhoneNumber : payloads.PhoneNumber,
	}

	err = s.vendorRepository.Update(vendorData, id)
	if err != nil {
		return dto.VendorResponse{}, err
	}

	vendorData, err = s.vendorRepository.GetById(id)
	if err != nil {
		return dto.VendorResponse{}, err
	}

	vendorResponse := dto.ConvertToVendorResponse(vendorData)
	
	return vendorResponse, nil
}

func (s *vendorUsecase) Delete(id int) error {
	_, err := s.vendorRepository.GetById(id)
	if err != nil {
		return err
	}

	err = s.vendorRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
