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

	//разрешение фронту обращаться к этому серверу
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	//

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, "все супер") })
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run() //на порту 8080
}
