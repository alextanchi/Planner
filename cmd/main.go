package main

import (
	"Planner"
	"Planner/pkg/handler"
	"log"
)

// комментарий
func main() {
	handlers := new(handler.Handler)                               //создаем экземпляр хендлер из нашего пакета pkg.handler
	srv := new(planner.Server)                                     // инициалиируем экзепляр сервера
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil { //если во время работы метод ран возвращает ошибку,
		// тогда вызываем метод Fatal, который выведет ошибку и выйдет из приложения
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
