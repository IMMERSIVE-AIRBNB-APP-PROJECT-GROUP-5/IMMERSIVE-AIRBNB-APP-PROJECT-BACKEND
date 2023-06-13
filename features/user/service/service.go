package service

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/AIRBNBAPP/features/user"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (user.Core, string, error) {
	// Mengatur validator
	validate := validator.New()
	Login := user.Login{
		Email:    email,
		Password: password,
	}
	errValidate := validate.Struct(Login)
	if errValidate != nil {
		return user.Core{}, "", errValidate
	}

	// Cek apakah pengguna mengirimkan data kosong untuk semua bidang
	if Login.Email == "" || Login.Password == "" {
		return user.Core{}, "", errors.New("Semua data harus diisi")
	}

	// Validasi email harus format email
	emailFormat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if Login.Email != "" && !emailFormat.MatchString(Login.Email) {
		return user.Core{}, "", errors.New("Format email tidak valid")
	}

	// Validasi panjang password minimal 8 karakter
	if Login.Password != "" && len(Login.Password) < 8 {
		return user.Core{}, "", errors.New("Password harus memiliki panjang minimal 8 karakter")
	}

	// Validasi password kombinasi huruf dan angka
	if Login.Password != "" {
		hasLetter := false
		hasDigit := false
		for _, ch := range Login.Password {
			if unicode.IsLetter(ch) {
				hasLetter = true
			} else if unicode.IsDigit(ch) {
				hasDigit = true
			}
		}
		if !hasLetter || !hasDigit {
			return user.Core{}, "", errors.New("Password harus kombinasi huruf dan angka")
		}
	}

	// Lanjutkan dengan proses login
	dataLogin, token, errValidate := service.userData.Login(email, password)
	return dataLogin, token, errValidate
}

// CreateUser implements user.UserServiceInterface.
func (service *userService) CreateUser(userInput user.Core) error {
	// Mengatur ulang validator
	validate := validator.New()
	Register := user.Register{
		User_name: userInput.User_name,
		Email:     userInput.Email,
		Password:  userInput.Password,
		Phone:     userInput.Phone,
	}

	errValidate := validate.Struct(Register)
	if errValidate != nil {
		return errValidate
	}

	// Cek apakah pengguna mengirimkan data kosong untuk semua bidang
	if userInput.User_name == "" || userInput.Email == "" || userInput.Password == "" || userInput.Phone == "" {
		return errors.New("Semua data harus diisi")
	}

	// Validasi email harus format email
	emailFormat := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if userInput.Email != "" && !emailFormat.MatchString(userInput.Email) {
		return errors.New("Format email tidak valid")
	}

	// Validasi panjang password minimal 8 karakter
	if userInput.Password != "" && len(userInput.Password) < 8 {
		return errors.New("Password harus memiliki panjang minimal 8 karakter")
	}

	// Validasi password kombinasi huruf dan angka
	if userInput.Password != "" {
		hasLetter := false
		hasDigit := false
		for _, ch := range userInput.Password {
			if unicode.IsLetter(ch) {
				hasLetter = true
			} else if unicode.IsDigit(ch) {
				hasDigit = true
			}
		}
		if !hasLetter || !hasDigit {
			return errors.New("Password harus kombinasi huruf dan angka")
		}
	}

	// Validasi nama hanya huruf
	if userInput.User_name != "" {
		for _, ch := range userInput.User_name {
			if unicode.IsDigit(ch) {
				return errors.New("Nama harus huruf")
			}
		}
	}

	// Validasi nomor telepon hanya angka
	if userInput.Phone != "" {
		for _, ch := range userInput.Phone {
			if !unicode.IsDigit(ch) {
				return errors.New("Phone harus angka")
			}
		}
	}

	// Validasi nomor telepon hanya angka dan minimal 10 digit
	if userInput.Phone != "" {
		digitCount := 0
		for _, ch := range userInput.Phone {
			if unicode.IsDigit(ch) {
				digitCount++
			}
		}
		if digitCount < 10 {
			return errors.New("Phone minimal 10 digit angka")
		}
	}
	errRegister := service.userData.CreateUser(userInput)
	return errRegister
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
