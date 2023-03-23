package response

import (
	"be-sharing-vision/app/domain"
)

type PostResponse struct {
	Id       int    `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	Category string `json:"category" form:"catgory"`
	Status   string `json:"status" form:"status"`
}

func PostFromDomain(domain domain.PostDomain) PostResponse {
	return PostResponse{
		Id:       domain.Id,
		Title:    domain.Title,
		Content:  domain.Content,
		Category: domain.Category,
		Status:   domain.Status,
	}
}
