package request

type LanguageCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
