package auth

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"

	"github.com/greetinc/greet-auth-srv/util"
)

func (u *authService) SignupDetail(req dto.RegisterDetailRequest) (dto.RegisterDetailResponse, error) {

	profileID := generateUniqueProfileID(req.FullName)

	user := dto.RegisterDetailRequest{
		ID:        util.GenerateRandomString(),
		FullName:  req.FullName,
		UserID:    req.UserID,
		ProfileID: profileID,
		Age:       req.Age,
		Gender:    req.Gender,
	}

	createdUser, err := u.Repo.SignupDetail(user)
	if err != nil {
		return dto.RegisterDetailResponse{}, err
	}

	response := dto.RegisterDetailResponse{
		ID:        createdUser.ID,
		FullName:  createdUser.FullName,
		ProfileID: createdUser.ProfileID,
		UserID:    user.UserID,
		Age:       createdUser.Age,
		Gender:    createdUser.Gender,
	}

	return response, nil
}

func generateUniqueProfileID(fullName string) string {
	// Extract the first and last characters of the full name
	firstChar := string(fullName[0])
	// Generate the current time in the format "150405" (HHMMSS)
	currentTime := time.Now().Format("150405")
	// Extract 2 characters from the end of the full name
	lastTwoChars := fullName[len(fullName)-2:]
	// Generate alphanumeric random string with a specified length (e.g., 4)
	randomDigits := generateRandomAlphaNumeric(4)

	// Generate a unique profile ID by combining the extracted characters, current time, and random digits
	profileID := fmt.Sprintf("%s%s%s%s", firstChar, currentTime, lastTwoChars, randomDigits)

	return profileID
}

// generateRandomAlphaNumeric generates a random alphanumeric string of a specified length
func generateRandomAlphaNumeric(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var result strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}

	return result.String()
}

func GenerateRandomNumeric(length int) string {
	const chars = "0123456789"

	var result strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}

	return result.String()
}
