package utils

import (
	"fmt"

	// "ginchat/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Config app inited")
}

func InitMySQL() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	fmt.Println("MySQL inited")
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)
}
