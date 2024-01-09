package handler

import "github.com/gin-gonic/gin"

// создаем пустые обработчики и присваиаем к маршрутам
// описываем пустые хендлреры для эндпоинтов регистрации и авторизации
// хендлер во фрейморке gin это функция которая должна иметь в качестве параметра указатель на объект gin.Context
func (h *Handler) signUp(c *gin.Context) {

}

func (h *Handler) signIn(c *gin.Context) {

}
