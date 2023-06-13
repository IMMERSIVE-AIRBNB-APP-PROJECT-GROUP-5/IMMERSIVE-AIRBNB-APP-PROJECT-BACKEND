package handler

type UserStatus string

const (
	Reservant UserStatus = "reservant"
	Hosters   UserStatus = "hosters"
)

type UserResponse struct {
	User_name string `json:"nama" form:"nama"`
	Email     string `json:"email" form:"email"`
	Phone     string `json:"phone" form:"phone"`
	Status    UserStatus
}
