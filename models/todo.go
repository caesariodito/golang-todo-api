package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          string
	Task        string
	Description string
	IsFinished  bool
}
