package request

type PostRequest struct {
	Title    string `json:"title" form:"title" validate:"required,max=20"`
	Content  string `json:"content" form:"content" validate:"required,max=20"`
	Category string `json:"category" form:"category" validate:"required,min=3"`
	Status   string `json:"status" form:"status" validate:"required"`
}
