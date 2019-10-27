package main

import "github.com/gin-gonic/gin"

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "pong",
			"code": 0,
		})
	}

}

//submit new task
func newTask() gin.HandlerFunc {
	return func(c *gin.Context) {

		//todo: parse request,get task parameter

		//todo: write to db

		//return
		c.JSON(200,gin.H{
			"message": "post ok",
			"code": 0,
		})

	}
}

func main() {
	r := gin.Default()

	// return some thing
	r.GET("/ping", ping())

	// "POST" -- receive http request
	r.POST("/tasks",newTask())

	// todo: add all other interface

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
