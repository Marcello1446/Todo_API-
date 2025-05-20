package utils

import (
	"Todo/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate = validator.New()

func ValidateTask(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task

		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := validate.Struct(&task); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		c.Set("task", &task)
		return next(c)
	}
}
