package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // инициализируем ротер

	//объявляем методы, сгруппировав их по маршрутам
	auth := router.Group("/auth") //группа auth
	{
		auth.POST("/sign-up", h.signUp) //эндпоинт
		auth.POST("/sign-in", h.signIn) //эндпоинт
	}
	api := router.Group("/api") //группа api будет использоваться для эндопинтов работ со списками и их задачами
	{
		lists := api.Group("/lists") //внутри группы api создаем еще одну группу lists,
		// которая будет использована для работы со списками
		{
			lists.POST("/", h.createList)      // создание списка
			lists.GET("/", h.getAllLists)      // получение всех списков
			lists.GET("/:id", h.getListById)   // получение всех списков по id
			lists.PUT("/:id", h.updateList)    // редактирование списка
			lists.DELETE("/:id", h.deleteList) // удаление списка
			//используя двоеточие в маршруте эндпоинта мы указываем,
			//что тут может быть любое значение, к которому мы можем обратиться по имени параметра id (фишка gin)
			items := lists.Group(":id/items") //группа для задач списка, делаем аналогично написанному выше
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)

			}

		}
		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
