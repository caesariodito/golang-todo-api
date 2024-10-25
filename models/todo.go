package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model  `json:"-"`
	ID          string `json:"id" gorm:"primaryKey"`
	Task        string `json:"task"`
	Description string `json:"description"`
	IsFinished  bool   `json:"is_finished"`
}
