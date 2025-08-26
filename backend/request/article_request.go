package request

type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=Publish Draft Trash"`
}

type UpdateArticleRequest struct {
	Title    string `json:"title" validate:"omitempty,min=20"`
	Content  string `json:"content" validate:"omitempty,min=200"`
	Category string `json:"category" validate:"omitempty,min=3"`
	Status   string `json:"status" validate:"omitempty,oneof=Publish Draft Trash"`
}
