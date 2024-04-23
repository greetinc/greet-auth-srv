package services

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"

	"github.com/greetinc/greet-auth-srv/entity"
)

func (b *countryService) GetById(req dto.CountryRequest) (entity.Country, error) {
	countrys, err := b.CountryR.GetById(req)
	if err != nil {
		return entity.Country{}, err
	}

	return countrys, nil
}
