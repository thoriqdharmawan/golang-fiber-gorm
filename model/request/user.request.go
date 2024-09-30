package request

type UserCreateRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Languages []uint `json:"languages"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type UserSetLanguageRequest struct {
	UserId     uint `json:"user_id" validate:"required"`
	LanguageId uint `json:"language_id" validate:"required"`
}
