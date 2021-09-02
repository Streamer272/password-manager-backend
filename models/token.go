package models

type Token struct {
	Id      uint  `gorm:"primaryKey" json:"id"`
	Expires int64 `gorm:"notNull" json:"expires"`
	UserId  uint  `gorm:"notNull" json:"user_id"`
	User    User  `gorm:"foreignKey:UserId;references:id" json:"-"`
}
