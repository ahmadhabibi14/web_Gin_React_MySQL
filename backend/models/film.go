package models

type Film struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Year  int    `json:"release_year"`
	Genre string `json:"genre"`
}
