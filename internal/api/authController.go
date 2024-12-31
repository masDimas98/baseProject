package api

import (
	"github.com/gin-gonic/gin"
	"goTest/internal/service"
)

// Login controller function of Login
func Login(c *gin.Context) {
	//access to auth service func LoginHandler
	data, err := service.LoginHandler(c)
	//error handler
	if err != nil {
		//if error happen returning errorResponse baseResponse
		c.JSON(err.ErrorCode, err)
	} else {
		//returning successResponse baseResponse
		c.JSON(data.SuccessCode, data)
	}
}

func Logout(c *gin.Context) {
	data, err := service.Logout(c)
	if err != nil {
		//if error happen returning errorResponse baseResponse
		c.JSON(err.ErrorCode, err)
	} else {
		//returning successResponse baseResponse
		c.JSON(data.SuccessCode, data)
	}
}
