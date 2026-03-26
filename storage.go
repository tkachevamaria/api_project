package main

import (
	"errors"
	"sync" //для mutex
)

type Employee struct {
	ID int `json:"id"` //тег json:"id" указывает,
	// что при сериализации структуры в JSON, поле ID будет отображаться как "id"
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

type MemoryStorage struct { //реализует интерфейс Storage если у него есть все методы
	counter    int
	data       map[int]Employee
	sync.Mutex //встраивание!!!!!!! (типа если имя не объявить)
	// я смогу обращаться к методам
	// sync.Mutex напрямую через экземпляр MemoryStorage
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()
	e.ID = s.counter
	s.data[e.ID] = *e //звезда у имени объекта, значит берем сам объект а не ссылку
	s.counter++
	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock() //по дефолту в конце

	employee, ok := s.data[id]
	if !ok {
		return employee, errors.New("employee not found")
	}
	return employee, nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id) //встроенная функция удаления из мапы по ключу
	s.Unlock()
}
