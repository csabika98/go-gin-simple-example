package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initDb()        // Use the function from database.go to initialize the database
	DefineRoutes(r) // Use the function from routes.go to define routes
	r.Run()         // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
