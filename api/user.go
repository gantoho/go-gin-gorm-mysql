package api

import (
	"github.com/gin-gonic/gin"
	"gormtest/dao"
	"strconv"
	"time"
)

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "李四",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.SaveUser(user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var user dao.User
	user = dao.GetUserById(id)
	c.JSON(200, user)
}

func GetAll(c *gin.Context) {
	var users []dao.User
	users = dao.GetAllUser()
	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	dao.UpdateUser(id)
}
