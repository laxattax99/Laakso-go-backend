package models

type VisitorCounter struct {
	ID    uint `gorm:"primaryKey"`
	Count int
}
