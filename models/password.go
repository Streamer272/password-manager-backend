package models

type Password struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"notNull" json:"name"`
	Value  string `gorm:"notNull" json:"value"`
	UserId uint   `gorm:"notNUll" json:"user_id"`
	User   User   `gorm:"foreignKey:UserId;references:id" json:"-"`
}
