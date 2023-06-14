package handler

type UserRequest struct {
	User_name string `json:"nama" form:"nama"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	Status    string
}

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UploadRequest struct {
	Url string `form:"url"`
}
