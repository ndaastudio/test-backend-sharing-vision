package model

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title" gorm:"type:varchar(200)" validate:"required"`
	Content     string    `json:"content" gorm:"type:text" validate:"required"`
	Category    string    `json:"category" gorm:"type:varchar(100)" validate:"required"`
	Status      string    `json:"status" gorm:"type:enum('Publish','Draft','Trash')" validate:"required"`
	CreatedDate time.Time `json:"created_date" gorm:"type:datetime"`
	UpdatedDate time.Time `json:"updated_date" gorm:"type:datetime"`
}
