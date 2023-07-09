package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine{
	r := gin.Default();
	r.GET("/user",func(ctx *gin.Context) { 
		ctx.JSON(http.StatusOK, gin.H{ 
			"ID" : "success",
			"name" : "Hello Sahal",
		})
})
	

return r
}


func main()  {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}