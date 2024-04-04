package seeder

import (
	"greet-auth-srv/configs"
	"greet-auth-srv/entity"
)

func Users() {
	db := configs.InitDB()
	// now := time.Now()

	var users []entity.User
	var budi = entity.User{
		ID: "956f2014-f8ab-41e2-88c1-0c3871524665",
		// FullName: "Budi",
		Email:    "budi@gmail.com",
		Password: "$2a$10$k7oCB0eh840JtXnQ74OnUezmBuYLQnmdbXOLc3ztN9F7y/C8jHFE6", //12345678
	}
	users = append(users, budi)

	var annisa = entity.User{
		ID: "c3c9e0d3-b3cf-4e53-9d24-b38f16e6f419",
		// FullName: "Annisa",
		Email:    "annisa@gmail.com",
		Password: "$2a$10$k7oCB0eh840JtXnQ74OnUezmBuYLQnmdbXOLc3ztN9F7y/C8jHFE6", //12345678

	}
	users = append(users, annisa)

	if err := db.Create(&users).Error; err != nil {
		return
	}
}

func RunSeeder() {
	Users()
}
