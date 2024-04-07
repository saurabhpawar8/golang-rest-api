package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhpawar8/golang-rest-api/db"
	"github.com/saurabhpawar8/golang-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	// server.Use(corsMiddleware())
	server.Run("localhost:8080")
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 			return
// 		}

// 		c.Next()
// 	}
// }
