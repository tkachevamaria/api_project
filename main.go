package main

import (
	"net/http"

	"github.com/gin-gonic/gin" //фреймворк для создания веб-приложений на языке Go
)

func main() {

	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage) //внедрение зависимостей

	router := gin.Default() //объект типа *gin.Engine
	//создание нового экземпляра роутера с помощью функции Default(),
	// которая возвращает роутер с предопределенными средними программами
	// (middleware) для обработки запросов и ответов

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, "все супер") })
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run() //на порту 8080
}
