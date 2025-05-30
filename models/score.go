package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	PlayerID string
	Score    int
	Combo    int
	PlayTime float64
}
