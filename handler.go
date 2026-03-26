package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"` // tag json
}

//структура для возврата ошибки

type Handler struct { //структура, у которой должно быть хранилице "типа" storage
	storage Storage
}

func NewHandler(storage Storage) *Handler { //создаем новый хэндлер, выходное значение по ссылке естессна
	//storage не передается по ссылке, потому что интерфейсы уже ссылочные
	return &Handler{storage: storage} //по ссылке
}

// -------------------
// хендлеры под 4 роута:
// Это функция, которую вызывает роутер Gin, когда приходит http запрос
func (h *Handler) CreateEmployee(c *gin.Context) {
	//Объект*gin.Context хранит в себе всю информацию про входящий запрос и умеет
	//также записывать ответы
	var employee Employee
	if err := c.BindJSON(&employee); err != nil { //& потому что передается сыылка на объект!
		//BindJSON пытается распарсить тело запроса в employee
		//и заполнить employee данными
		fmt.Printf("Failed to bind employee: %s\n", err.Error()) //вывод в консоль
		c.JSON(http.StatusBadRequest, ErrorResponse{             //надо разобраться как именно работает, я не понимаю
			Message: err.Error(),
		})
		//c.JSON - автоматически преобразует структуру в JSON и СРАЗУ отправляет клиенту
		//первым агрументом c.JSON требует статус-код (особенность gin)
		return
	}

	h.storage.Insert(&employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	//ошибка получения id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	//проблема привязки сотрудника (нового)
	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Update(id, employee) //апдейтим хранилище

	c.JSON(http.StatusOK, "все заебись, сотрудник обновлен") //вывод айдишника сотрудника не понимаю зачем правда

}

func (h *Handler) GetEmployee(c *gin.Context) {

	//ошибка преобразования типа
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	//ошибка получения сотрудника
	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)

}

func (h *Handler) DeleteEmployee(c *gin.Context) {

	//ошибка преобразования
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Delete(id)

	c.String(http.StatusOK, "employee deleted") //типа строку вывести а не структуру
}
