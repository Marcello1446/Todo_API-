package handler

import (
	"Todo/database"
	"Todo/models"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"unicode"
)

func GetAllTasks(c echo.Context) error {
	var tasks []models.Task

	if err := database.DB.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, tasks)
}

func CreateTask(c echo.Context) error {
	task := c.Get("task").(*models.Task)

	if err := database.DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := database.DB.Save(&task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.DB.Delete(&task, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, task)
}

func CreateTwoTasks(tx *gorm.DB, task1, task2 *models.Task) error {
	prepareAndValidate := func(t *models.Task) error {
		if len(t.Title) < 6 {
			return fmt.Errorf("задача \"%s\" — заголовок слишком короткий", t.Title)
		}

		runes := []rune(t.Title)
		if !unicode.IsUpper(runes[0]) {
			runes[0] = unicode.ToUpper(runes[0])
			t.Title = string(runes)
		}

		var count int64
		if err := tx.Model(&models.Task{}).Where("title = ?", t.Title).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("задача \"%s\" — дубликат", t.Title)
		}

		return nil
	}

	return tx.Transaction(func(tx2 *gorm.DB) error {
		if err := prepareAndValidate(task1); err != nil {
			return fmt.Errorf("первая задача: %w", err)
		}
		if err := prepareAndValidate(task2); err != nil {
			return fmt.Errorf("вторая задача: %w", err)
		}

		if task1.Title == task2.Title {
			return errors.New("заголовки задач не должны совпадать")
		}

		if err := tx2.Create(task1).Error; err != nil {
			return err
		}
		if err := tx2.Create(task2).Error; err != nil {
			return err
		}

		return nil
	})
}
