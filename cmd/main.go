package main

import (
	"Planner"
	"Planner/pkg/handler"
	"Planner/pkg/repository"
	"Planner/pkg/service"
	"log"
)

// основная логика здесь
func main() {
	repos := repository.NewRepository() //объявлем все зависимости в нужном порядке
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(planner.Server)                                     // инициалиируем экзепляр сервера
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil { //если во время работы метод ран возвращает ошибку,
		// тогда вызываем метод Fatal, который выведет ошибку и выйдет из приложения
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
