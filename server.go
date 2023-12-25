package planner

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error { //запуск
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, //1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe() //слушаем все входящие запросы для дальнейшей обработки
}
func (s *Server) Shutdown(ctx context.Context) error { // остановка сервера
	return s.httpServer.Shutdown(ctx)
}
