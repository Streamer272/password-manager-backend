package models

type Password struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	UserId uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserId;references:id" json:"-"`
}
