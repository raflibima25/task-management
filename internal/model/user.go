package model

import "gorm.io/gorm"

type User struct {
	ID        string `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email     string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Password  string `json:"-" gorm:"type:varchar(255);not null"`
	Tasks     []Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt string `json:"deleted_at" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (User) TableName() string {
	return "users"
}
