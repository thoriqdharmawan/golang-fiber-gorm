package request

type PostCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	UserID uint   `json:"user_id" validate:"required"`
}
