package service

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags index
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")

	// c.JSON(200, gin.H{
	// 	"message": "welcome",
	// })
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")

	// c.JSON(200, gin.H{
	// 	"message": "welcome",
	// })
}
