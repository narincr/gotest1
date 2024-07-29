package main

import (
	"Test1/api"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"os"
	"time"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // เปลี่ยน
	router := gin.Default()
	// Set up CORS middleware options
	config := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     false,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to Ythe router (works on groups too)
	router.Use(cors.Middleware(config))
	api.Setup(router)
	err := router.Run(":8888")
	if err != nil {
		fmt.Println("Error[1]:", err)
	}
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("Running on heroku using random PORT")
		err := router.Run()
		if err != nil {
			fmt.Println("Error[2]:", err)
		}
	} else {
		fmt.Println("Environment Port " + port)
		err := router.Run(":" + port)
		if err != nil {
			fmt.Println("Error[3]:", err)
		}
	}
}
