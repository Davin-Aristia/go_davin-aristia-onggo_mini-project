package usecase

import (

	"go-mini-project/dto"
	"go-mini-project/model"
	"go-mini-project/repository"
)

type VendorUsecase interface {
	Get(name string) ([]model.Vendor, error)
	GetById(id int) (model.Vendor, error)
	Create(payloads dto.VendorRequest) (model.Vendor, error)
	Update(payloads dto.VendorRequest, id int) (model.Vendor, error)
	Delete(id int) error
}

type vendorUsecase struct {
	vendorRepository repository.VendorRepository
}

func NewVendorUsecase(vendorRepo repository.VendorRepository) *vendorUsecase {
	return &vendorUsecase{vendorRepository: vendorRepo}
}

func (s *vendorUsecase) Get(name string) ([]model.Vendor, error) {
	vendorData, err := s.vendorRepository.Get(name)
	if err != nil {
		return []model.Vendor{}, err
	}

	return vendorData, nil
}

func (s *vendorUsecase) GetById(id int) (model.Vendor, error) {
	vendorData, err := s.vendorRepository.GetById(id)
	if err != nil {
		return model.Vendor{}, err
	}

	return vendorData, nil
}

func (s *vendorUsecase) Create(payloads dto.VendorRequest) (model.Vendor, error) {
	vendorData := model.Vendor{
		Name : payloads.Name,
		Address : payloads.Address,
		Email : payloads.Email,
		PhoneNumber : payloads.PhoneNumber,
	}

	vendorData, err := s.vendorRepository.Create(vendorData)
	if err != nil {
		return model.Vendor{}, err
	}
	return vendorData, nil
}

func (s *vendorUsecase) Update(payloads dto.VendorRequest, id int) (model.Vendor, error) {
	_, err := s.vendorRepository.GetById(id)
	if err != nil {
		return model.Vendor{}, err
	}
	
	vendorData := model.Vendor{
		Name : payloads.Name,
		Address : payloads.Address,
		Email : payloads.Email,
		PhoneNumber : payloads.PhoneNumber,
	}

	err = s.vendorRepository.Update(vendorData, id)
	if err != nil {
		return model.Vendor{}, err
	}

	vendorData, err = s.vendorRepository.GetById(id)
	if err != nil {
		return model.Vendor{}, err
	}
	
	return vendorData, nil
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
