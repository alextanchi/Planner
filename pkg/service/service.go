package service

import "Planner/pkg/repository"

//интерфейсы называем исходя из их доменной зоны (т.е участка бизнес логики приложения за которую они отвечают)

type Autorization interface {
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Autorization
	TodoList
	TodoItem
}

// объявляем конструктор

func NewService(repos *repository.Repository) *Service { //сервисы будут обращаться к БД,
	// поэтому объявляем указатель на структуру репозитория в качестве аргумента
	return &Service{}
}
