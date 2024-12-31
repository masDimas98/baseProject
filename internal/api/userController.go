package api

import (
	"github.com/gin-gonic/gin"
	"goTest/internal/service"
	"log"
)

func GetUsers(c *gin.Context) {
	err := service.VerifyRequest(c)
	if err != nil {
		c.JSON(err.ErrorCode, err)
	} else {
		data, err := service.GetAllUsers(c)
		if err != nil {
			c.JSON(err.ErrorCode, err)
		} else {
			c.JSON(data.SuccessCode, data)
		}
	}
}

func GetUser(c *gin.Context) {
	err := service.VerifyRequest(c)
	if err != nil {
		c.JSON(err.ErrorCode, err)
	} else {
		data, err := service.GetUser(c)
		if err != nil {
			log.Printf("data : %v", err)
			c.JSON(err.ErrorCode, err)
		} else {
			c.JSON(data.SuccessCode, data)
		}
	}
}

func UpdateUser(c *gin.Context) {
	err := service.VerifyRequest(c)
	if err != nil {
		c.JSON(err.ErrorCode, err)
	} else {
		data, err := service.UpdateUser(c)
		if err != nil {
			log.Printf("data : %v", err)
			c.JSON(err.ErrorCode, err)
		} else {
			c.JSON(data.SuccessCode, data)
		}
	}
}

func CreateUser(c *gin.Context) {
	data, err := service.CreateUser(c)
	if err != nil {
		c.JSON(err.ErrorCode, err)
	} else {
		c.JSON(data.SuccessCode, data)
	}
}

func DeleteUser(c *gin.Context) {
	err := service.VerifyRequest(c)
	if err != nil {
		c.JSON(err.ErrorCode, err)
	} else {
		data, err := service.DeleteUser(c)
		if err != nil {
			c.JSON(err.ErrorCode, err)
		} else {
			c.JSON(data.SuccessCode, data)
		}
	}
}
