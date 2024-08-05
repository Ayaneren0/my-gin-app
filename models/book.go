package models

type Book struct {
	ID     string `gorm:"id" binding:"required"`
	Title  string `gorm:"title" binding:"required"`
	Author string `gorm:"author" binding:"required"`
}
