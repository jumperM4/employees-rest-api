package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)
	fmt.Println(memoryStorage, handler)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run(":8000")
}
