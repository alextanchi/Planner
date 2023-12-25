package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // инициализируем ротер

	//объявляем методы, сгруппировав их по маршрутам
	auth := router.Group("/auth") //группа auth
	{
		auth.POST("/sign-up") //эндпоинт
		auth.POST("/sign-in") //эндпоинт
	}
	api := router.Group("/api") //группа api будет использоваться для эндопинтов работ со списками и их задачами
	{
		lists := api.Group("/lists") //внутри группы api создаем еще одну группу lists,
		// которая будет использована для работы со списками
		{
			lists.POST("/")      // создание списка
			lists.GET("/")       // получение всех списков
			lists.GET("/:id")    // получение всех списков по id
			lists.PUT("/:id")    // редактирование списка
			lists.DELETE("/:id") // удаление списка
			//используя двоеточие в маршруте эндпоинта мы указываем,
			//что тут может быть любое значение, к которому мы можем обратиться по имени параметра id (фишка gin)
			items := lists.Group(":id/items") //группа для задач списка
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				lists.PUT("/:item_id")
				lists.DELETE("/:item_id")
			}
		}
	}
	return router
}
