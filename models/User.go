package models

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"-"`
	Username  string    `gorm:"unique;not null;varchar(255)" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Photo     []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `gorm:"not_null;autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"not_null;autoUpdateTime" json:"-"`
}
