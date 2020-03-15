package main

import (
	"day_frame/db_config"
	"day_frame/tasks"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db_config.Connect()
	// Setup router
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./ui/build", true)))

	// Setup router for the api
	api := router.Group("/api")

	//Api routes
	api.GET("/tasks/:userId", tasks.GetTasks)
	api.POST("/tasks", tasks.PostTask)
	api.PUT("/tasks/:taskId", tasks.UpdateTask)
	api.DELETE("/tasks/:taskId", tasks.DeleteTask)

	//Start and run server
	log.Fatal(router.Run(":5000"))

}
