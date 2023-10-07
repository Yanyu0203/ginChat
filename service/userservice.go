package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Summary user list
// @Tags User
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{
		"code":    0,
		"message": "User list",
		"data":    data,
	})
}

// CreateUser
// @Summary create user
// @Tags User
// @param name query string false "userName"
// @param password query string false "password"
// @param repassword query string false "repassword"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}

	// user.Name = c.Query("name")
	// password := c.Query("password")
	// repassword := c.Query("repassword")
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("Identity")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "Name or password is empty",
			"data":    user,
		})
		return
	}
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "This name has been used",
			"data":    user,
		})
		return
	}

	if password != repassword {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "password dont match",
			"data":    user,
		})
		return
	}

	// user.PassWord = passWord
	user.Salt = salt
	user.PassWord = utils.MakePassword(password, salt)
	user.LoginTime = time.Now()
	user.LogoutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Sign in success",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// @Summary User log in
// @Tags User
// @param name query string false "userName"
// @param password query string false "password"
// @Success 200 {string} json{"code", "message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}

	// name := c.Query("name")
	// passWord := c.Query("passWord")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "User not exist",
			"data":    data,
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //0, success; 1, fail
			"message": "Password not match",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)

	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    0, //0, success; 1, fail
		"message": "Log in success",
		"data":    data,
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
		"code":    0,
		"message": "Delete user success",
		"data":    user,
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
			"code":    -1,
			"message": "param no match",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Update user success",
			"data":    user,
		})
	}

}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "Serch friends success!",
		"data":    users,
	})
}
