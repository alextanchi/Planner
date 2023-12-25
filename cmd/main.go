package main

import (
	"Planner"
	"log"
)

// комментарий
func main() {
	srv := new(planner.Server)              // инициалиируем экзепляр сервера
	if err := srv.Run("8000"); err != nil { //если во время работы метод ран возвращает ошибку,
		// тогда вызываем метод Fatal, который выведет ошибку и выйдет из приложения
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
