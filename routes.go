package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Define the database connection
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "database/data.db")
	if err != nil {
		panic(err)
	}
}

func DefineRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "World",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")

		// Insert the data into the database
		_, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"email":    email,
		})
	})
	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer rows.Close()

		var users []gin.H
		for rows.Next() {
			var id int
			var username, email string
			err = rows.Scan(&id, &username, &email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			users = append(users, gin.H{
				"id":       id,
				"username": username,
				"email":    email,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})
}
