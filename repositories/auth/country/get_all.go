package repositories

import (
	"github.com/greetinc/greet-auth-srv/entity"
)

func (r *countryRepository) GetAll() ([]entity.Country, error) {
	var countrys []entity.Country

	err := r.DB.Model(entity.Country{}).Find(&countrys).Error
	if err != nil {
		return nil, err
	}

	return countrys, nil
}
