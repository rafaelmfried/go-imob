package entities

import (
	"time"

	uuid "github.com/google/uuid"
)

type Property struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Address 	 string    `gorm:"size:300;not null" json:"address"`
	CreatedAt	 time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt	 time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	OwnerID     uuid.UUID      `gorm:"type:uuid; not null" json:"owner_id"`
}