package models

import "time"

type Destination struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	Category  string `gorm:"type:varchar(100)"` 
	Location  string `gorm:"type:varchar(255)"` 
	Latitude  float64
	Longitude float64
	Rating    float64 `gorm:"type:decimal(3,2)"`
	CreatedAt time.Time
}