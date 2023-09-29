package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags userBasic
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}
