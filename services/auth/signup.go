package auth

import (
	"crypto/tls"
	dto "greet-auth-srv/dto/auth"
	res "greet-auth-srv/util/response"
	"strings"

	"greet-auth-srv/util"

	"gopkg.in/gomail.v2"
)

func (u *authService) Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {

	encryp := util.EncryptPassword(&req)
	if encryp != nil {
		return dto.AuthRegisterResponse{}, encryp
	}
	if !util.IsValidEmail(req.Email) {
		return dto.AuthRegisterResponse{}, res.ErrorBuilder(&res.ErrorConstant.RegisterMail, nil)
	}

	user := dto.AuthRegisterRequest{
		ID:       util.GenerateRandomString(),
		Otp:      GenerateRandomNumeric(4),
		Whatsapp: req.Whatsapp,
		Email:    strings.ToLower(req.Email),
		Password: req.Password,
		Token:    util.GenerateRandomString(),
	}

	createdUser, err := u.Repo.Signup(user)
	if err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	if err := util.Mailtrap(user.Email, user.Otp); err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	response := dto.AuthRegisterResponse{
		ID:       createdUser.ID,
		Whatsapp: createdUser.Whatsapp,
		Email:    createdUser.Email,
		Password: createdUser.Password,
		Token:    createdUser.Token,
	}

	return response, nil
}

func (u *authService) sendMail(to, verificationToken string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "aseprayana95@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Test Email")
	mailer.SetBody("text/html", "Hello, this is a test email from MailHog.")

	dialer := gomail.NewDialer("localhost", 1025, "", "")

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

// sendVerificationEmail mengirim email verifikasi ke alamat email.
func (u *authService) sendVerificationEmail(to, verificationToken string) error {
	// Konfigurasi pengaturan email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "aseprayana95@gmail.com") // Ganti dengan alamat email Gmail pengirim
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Account Verification")
	mailer.SetBody("text/html", "Click the following link to verify your account: "+
		util.GetVerificationLink(verificationToken))

	// Konfigurasi pengaturan koneksi email untuk Gmail
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aseprayana95@gmail.com", "tybm gndz imkq deev")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Hanya gunakan ini dalam pengembangan, tidak aman untuk produksi

	// Kirim email
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
