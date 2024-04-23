package repositories

import (
	"github.com/greetinc/greet-auth-srv/entity"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
)

func (b *countryRepository) GetById(req dto.CountryRequest) (entity.Country, error) {
	tr := entity.Country{
		Country: req.Country,
	}

	if err := b.DB.Where("country = ?", tr.Country).Take(&tr).Error; err != nil {
		return entity.Country{}, err
	}

	return tr, nil
}
