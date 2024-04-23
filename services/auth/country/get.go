package services

import (
	"github.com/greetinc/greet-auth-srv/entity"
)

func (s *countryService) GetAll() ([]entity.Country, error) {
	// Fetch data from the repository layer
	data, err := s.CountryR.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
