package models

type Task struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"  validate:"required,min=6"`
	Completed bool   `json:"completed"`
}

type CreateTwoTasksRequest struct {
	Task1Title string `json:"task1_title"`
	Task2Title string `json:"task2_title"`
}
