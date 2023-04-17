package Model

type City_info_structs struct {
	ID         uint `gorm:"primaryKey"`
	City       string
	Population int
}
