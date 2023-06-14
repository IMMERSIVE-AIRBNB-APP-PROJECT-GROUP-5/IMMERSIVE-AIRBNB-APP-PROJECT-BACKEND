package user

import "time"

type UserStatus string

const (
	Reservant UserStatus = "reservant"
	Hosters   UserStatus = "hosters"
)

type Core struct {
	Id         uint
	User_name  string `json:"user_name" form:"user_name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Phone      string `json:"phone" form:"phone"`
	Status     UserStatus
	Url        string `json:"url" form:"url"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type Login struct {
	Email    string `json:"nama" form:"name"`
	Password string `json:"password" form:"password"`
}

type Register struct {
	User_name string `json:"nama" form:"nama"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
}

type UpdatedProfil struct {
	User_name string `json:"nama" form:"nama"`
	Phone     string `json:"phone" form:"phone"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type UploadPicture struct {
	Url string `json:"url" form:"url"`
}

type UserDataInterface interface {
	Register(userInput Core) error
	Login(email, password string) (Core, string, error)
	Profil(id int) ([]Core, error)
	UpdatedProfil(id int, userInput Core) error
	DeleteAccount(id int) error
	ValidateHoster(id int, userInput Core) error
	// GetAllUser(keyword string) ([]Core, error)
	// GetUserByID(userID uint) (*Core, error)
	// GetRoleByID(userID int) (UserRole, error)
	// Update(userID int, updatedUser Core) error
	// Delete(userID int) error

}

type UserServiceInterface interface {
	Register(userInput Core) error
	Login(email, password string) (Core, string, error)
	Profil(id int) ([]Core, error)
	UpdatedProfil(id int, userInput Core) error
	DeleteAccount(id int) error
	ValidateHoster(id int, userInput Core) error
	// ValidateHoster(id int, url string) error
	// GetRoleByID(userID int) (UserRole, error)
	// GetAllUser(keyword string) ([]Core, error)
	// Update(userID int, updatedUser Core, loggedInUserID int) error
	// Delete(userID int, loggedInUserID int) error

}
