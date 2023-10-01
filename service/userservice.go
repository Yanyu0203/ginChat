package service

import (
	"fmt"
	"ginchat/models"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary user list
// @Tags User
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary create user
// @Tags User
// @param name query string false "userName"
// @param passWord query string false "passWord"
// @param repassWord query string false "repassWord"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}

	user.Name = c.Query("name")
	passWord := c.Query("passWord")
	repassWord := c.Query("repassWord")
	if passWord != repassWord {
		c.JSON(-1, gin.H{
			"message": "password dont match",
		})
		return
	}
	user.PassWord = passWord
	user.LoginTime = time.Now()
	user.LogoutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "Sign in success",
	})

}

// DeleteUser
// @Summary delete user
// @Tags User
// @param id query string false "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}

	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "Delete user success",
	})
}

// UpdateUser
// @Summary update user
// @Tags User
// @param id formData string false "id"
// @param name formData string false "name"
// @param passWord formData string false "passWord"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}

	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("passWord")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "param no match",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "Update user success",
		})
	}

}
