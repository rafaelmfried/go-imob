package entities

import (
	time "time"

	uuid "github.com/google/uuid"
	gorm "gorm.io/gorm"
)

type User struct {
	ID 			uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuidv7()" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
	Email   	string    `gorm:"size:150;uniqueIndex;not null" json:"email"`
	CreatedAt 	time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Properties []Property `json:"properties,omitempty" gorm:"foreignKey:OwnerID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.ID == uuid.Nil {
        u.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}