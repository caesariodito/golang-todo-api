package database

import (
	"fmt"
	"golang-todo-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func _Seeder(db *gorm.DB) {
	SeedTodos(db)
}

func SeedTodos(db *gorm.DB) {
	// Check if there are already todos in the database
	var count int64
	db.Model(&models.Todo{}).Count(&count)
	if count > 0 {
		fmt.Println("Todos already seeded, skipping seeder.")
		return
	}

	todos := []models.Todo{
		{
			ID:          uuid.New().String(),
			Task:        "Learn Golang",
			Description: "Complete Go struct and GORM tutorial",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Build Web App",
			Description: "Develop a web application backend",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Study Concurrency",
			Description: "Understand Go routines and channels",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Integrate GORM",
			Description: "Set up database ORM for the app",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Write Unit Tests",
			Description: "Implement tests for all functions",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Explore Go Modules",
			Description: "Learn about dependency management",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Optimize Queries",
			Description: "Analyze and improve SQL queries",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Create API Endpoints",
			Description: "Set up RESTful endpoints for the app",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Design User Interface",
			Description: "Draft the frontend for the application",
			IsFinished:  false,
		},
		{
			ID:          uuid.New().String(),
			Task:        "Deploy Application",
			Description: "Deploy the app to a cloud provider",
			IsFinished:  false,
		},
	}

	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			fmt.Println("Failed to seed todo:", err)
		} else {
			fmt.Println("Seeded todo:", todo.Task)
		}
	}
}
