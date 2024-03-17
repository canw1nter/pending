package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"pending/im/server/common"
	"pending/im/server/model"
)

type loginRequest struct {
	Name     string
	Password string
}

func loginHandler(c *gin.Context) {
	var params loginRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Printf("IP:%s request login failed! Couldn't bind params, err: %s\n", c.ClientIP(), err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user model.User
	result := model.MySQLDB.Model(&model.User{}).
		Where("name = ? AND password = ?", params.Name, params.Password).
		First(&user)
	if result.RowsAffected == 0 {
		log.Printf("User: %s login failed, Password: %s\n", params.Name, params.Password)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "Login failed",
			"data": nil,
		})
		return
	}

	token, err := common.GenerateUserToken(user.UUID, user.Name)
	if err != nil {
		log.Printf("Generate user %s's token failed! err: %s\n", user.Name, err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "Can't get a token",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Login successfully",
		"data": gin.H{
			"token": token,
		},
	})
}
