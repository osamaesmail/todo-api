package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type Todo struct {
	ID  uuid.UUID   `json:"id"`
	Title string  `json:"title" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Completed bool  `json:"completed"`
}

func (todo *Todo) BeforeCreate(scope *gorm.Scope) error {
	u, err := uuid.NewV4()
	if err !=nil {
		return err
	}

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", u.String())
	return nil
}

func (todo *Todo) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
