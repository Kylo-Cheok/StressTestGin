package api

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/6/18 18:27
 * @Desc:   Grace under pressure
 */
import (
	"fmt"
	controller "ginDemoProject/Controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUser(c *gin.Context) {
	name := c.Param("username")
	controller.GetUser(name, c)
}

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println(fmt.Sprintf("username:%s,password:%s", username, password))
	controller.CreateUser(username, password, c)
}
