package routes

import (
	"JET_DEV_TASK/JetDev/api/controller"
	"JET_DEV_TASK/JetDev/infrastructure"
)

type BlogRoute struct {
	Controller controller.BlogController
	Handler    infrastructure.GinRouter
}

func NewBLogRoute(
	controller controller.BlogController,
	handler infrastructure.GinRouter,

) BlogRoute {
	return BlogRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (b BlogRoute) Setup() {
	blog := b.Handler.Gin.Group("/blog") //Router group
	{
		blog.GET("/", b.Controller.ListAll)
		blog.GET("/", b.Controller.ListComment)
		blog.GET("/", b.Controller.ListContent)
		blog.POST("/", b.Controller.AddArticle)

	}
}
