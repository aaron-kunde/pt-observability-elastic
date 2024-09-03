package db

type DataEntity struct {
	ID   uint `gorm:"primaryKey"`
	Data string
}
