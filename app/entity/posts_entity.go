package entity

import (
	"be-sharing-vision/app/domain"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id        int    "gorm:primaryKey"
	Title     string `gorm:"type:varchar(200)"`
	Content   string `gorm:"type:text"`
	Category  string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string `gorm:"type:enum('Publish', 'Draft', 'Trash')"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Post) ToDomain(data Post) domain.PostDomain {
	return domain.PostDomain{
		Id:       data.Id,
		Title:    data.Title,
		Content:  data.Content,
		Category: data.Category,
		Status:   data.Status,
	}
}

func (p *Post) ToListDomain(data []Post) []domain.PostDomain {
	var listDomain []domain.PostDomain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain(item))
	}
	return listDomain
}
