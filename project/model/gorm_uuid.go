package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GormModel generate uuid v4
type GormModel struct {
	ID uuid.UUID `gorm:"type:varchar(256)" json:"id"`
}

// BeforeCreate generate uuid
func (model *GormModel) BeforeCreate(_ *gorm.DB) (err error) {
	model.ID = uuid.New()
	return
}
