package main

import (
	"go-bioskop/db"
	"go-bioskop/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.POST("/bioskop", handlers.CreateBioskop)
	r.GET("/bioskop", handlers.GetAllBioskop)
	r.GET("/bioskop/:id", handlers.GetBioskopByID)
	r.PUT("/bioskop/:id", handlers.UpdateBioskop)
	r.DELETE("/bioskop/:id", handlers.DeleteBioskop)

	r.Run(":8080")
}
