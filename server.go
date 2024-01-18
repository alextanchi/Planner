package planner

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error { //запуск, добавили в метод run аргумент хендлер типа интерфейса хендлер
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,          //передали handler в объект http сервера
		MaxHeaderBytes: 1 << 20,          //1MB
		ReadTimeout:    10 * time.Second, //10 сек
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe() //слушаем все входящие запросы для дальнейшей обработки
}
func (s *Server) Shutdown(ctx context.Context) error { // остановка сервера
	return s.httpServer.Shutdown(ctx)
}
