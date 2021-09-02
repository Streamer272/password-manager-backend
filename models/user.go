package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;notNull" validate:"required" json:"username"`
	Password string `gorm:"notNull" json:"-"`
}
