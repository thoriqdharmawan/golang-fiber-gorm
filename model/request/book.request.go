package request

type BookCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author"`
	Cover  string `json:"cover"`
}
