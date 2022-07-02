package user

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/5/22 15:07
 * @Desc:   Grace under pressure
 */
import (
	db "ginDemoProject/Database"
	"ginDemoProject/Models"
)

func FindUser(userName string) (user Models.User, err error) {
	result := db.QueryMysql(userName)
	return result, nil
}

func CreateUser(username string, password string) (err error) {
	result := db.CreateMysql(username, password)
	return result
}
