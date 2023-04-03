package model

type Book struct {
	ID          int    `gorm:"type:int; primary key;"`
	Name        string `gorm:"type:varchar(255)"`
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Cost        int    `gorm:"type:int"`
}
