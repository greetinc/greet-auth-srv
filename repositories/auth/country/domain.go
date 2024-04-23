package repositories

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"

	"github.com/greetinc/greet-auth-srv/entity"

	"gorm.io/gorm"
)

type CountryRepository interface {
	GetAll() ([]entity.Country, error)
	GetById(req dto.CountryRequest) (entity.Country, error)
}

type countryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(DB *gorm.DB) CountryRepository {
	return &countryRepository{
		DB: DB,
	}
}
