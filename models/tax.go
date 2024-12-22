package models

import "time"

type Tax struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Value     int64     `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
