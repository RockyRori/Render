package models

type Game struct {
	ID    uint    `gorm:"primaryKey"`
	Title string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
}
