package main

import (
	"Planner"
	"log"
)

func main() {
	srv := new(planner.Server)              // инициалиируем экзепляр сервера
	if err := srv.Run("8000"); err != nil { //если во время работы метод ран возвращает ошибку то выводим ошибку и выходим
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
