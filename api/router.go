package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
	r.GET("/get/:id", GetUser)
	r.GET("/getall", GetAll)
	r.GET("/updateuser/:id", UpdateUser)
}
