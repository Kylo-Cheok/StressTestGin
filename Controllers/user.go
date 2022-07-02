package controller

import (
	"fmt"
	"ginDemoProject/Models"
	"ginDemoProject/Pkg/e"
	stress "ginDemoProject/Services/stress_test"
	"ginDemoProject/Services/user"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/6/12 19:29
 * @Desc:   controller file
 */

func GetUser(userName string, c *gin.Context) {
	r, err := user.FindUser(userName)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, nil)
		return
	}
	c.JSON(e.SUCCESS, r)
}

func CreateUser(username string, password string, c *gin.Context) {
	err := user.CreateUser(username, password)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
	}
	c.JSON(e.SUCCESS, username+": Create Success")
}

func StartStress(num int, count int, r *Models.Request, c *gin.Context) {
	stress.StartTest(num, count, r)
	c.JSON(e.SUCCESS, "The stress test is in progress, please be patient.")
}
