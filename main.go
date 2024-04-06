package main

import (

	//"github.com/gin-contrib/cors"
	"sqltest/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "sqltest/database"
)

func main() {
	// db, err := database.InitDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	router := gin.Default()

	controllers.ConnectDatabase()

	router.Use(func(c *gin.Context) {
		// c.Set("db", db)
		c.Next()
	})

	router.POST("/library", controllers.CreateLibrary)
	router.GET("/library", controllers.GetAllLibrary)
	router.POST("/book/create", controllers.CreateBook)
	router.GET("/book", controllers.GetAllBook)
	router.PUT("/update/book/:isbn", controllers.UpdateBookByISBN)
	router.DELETE("/book/:isbn", controllers.DeleteBookByISBN)
	router.POST("/user", controllers.CreateUser)
	router.POST("/user/login", controllers.LoginUser)
	router.PUT("/user/:id", controllers.UpdateUserByID)
	router.GET("/book/:isbn", controllers.GetBookByISBN)
	router.POST("/request", controllers.CreateRequest)
	router.GET("/getrequest", controllers.GetAllRequest)
	router.PUT("/request/:reqid", controllers.UpdateRequestByReqID)
	router.GET("/user/:email", controllers.GetAllUser)
	router.POST("/issue", controllers.CreateIssue)

	router.Run(":8888")
}
