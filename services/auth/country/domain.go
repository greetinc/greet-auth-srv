package services

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	repositories "github.com/greetinc/greet-auth-srv/repositories/auth/country"
	m "github.com/greetinc/greet-middlewares/middlewares"

	"github.com/greetinc/greet-auth-srv/entity"
)

type CountryService interface {
	GetAll() ([]entity.Country, error)
	GetById(req dto.CountryRequest) (entity.Country, error)
}

type countryService struct {
	CountryR repositories.CountryRepository
	jwt      m.JWTService
}

func NewCountryService(CountryR repositories.CountryRepository, jwtS m.JWTService) CountryService {
	return &countryService{
		CountryR: CountryR,
		jwt:      jwtS,
	}
}
