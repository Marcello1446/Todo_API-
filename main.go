package main

import (
	"Todo/database"
	"Todo/handler"
	"Todo/middlewares"
	"Todo/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	e := echo.New()

	e.Use(middlewares.LogMiddleware)
	e.Use(middleware.CORS())

	e.GET("/tasks", handler.GetAllTasks)
	e.POST("/task", handler.CreateTask, utils.ValidateTask)
	e.PATCH("/task/:id", handler.UpdateTask)
	e.DELETE("/task/:id", handler.DeleteTask)

	err := e.Start(":8080")

	if err != nil {
		log.Fatal("Cannot connect to a server")
	}
}
