package controllers

import (
	"net/http"
	"prakerja_b4/configs"
	"prakerja_b4/models"

	"github.com/labstack/echo/v4"
)

func InsertUserController(c echo.Context) error {
	var userInput models.User
	c.Bind(&userInput)

	result := configs.DB.Create(&userInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Succes",
		Data:    userInput,
	})
}

func GetUserController(c echo.Context) error {
	var users []models.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, users)
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Succes",
		Data:    users,
	})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := configs.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var userInput models.User
	c.Bind(&userInput)

	user.Name = userInput.Name
	user.UserName = userInput.UserName
	user.Password = userInput.Password

	updateResult := configs.DB.Save(&user)
	if updateResult.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := configs.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
