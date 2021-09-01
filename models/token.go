package models

import (
	"time"
)

type Token struct {
	Id      uint      `gorm:"primaryKey" json:"id"`
	Expires time.Time `gorm:"notNull" json:"expires"`
	UserId  uint      `gorm:"notNull" json:"user_id"`
	User    User      `gorm:"foreignKey:UserId;references:id" json:"-"`
}
