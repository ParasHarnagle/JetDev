package main

import (
	"JET_DEV_TASK/JetDev/api/controller"
	"JET_DEV_TASK/JetDev/api/repository"
	"JET_DEV_TASK/JetDev/api/routes"
	"JET_DEV_TASK/JetDev/api/service"
	"JET_DEV_TASK/JetDev/infrastructure"
	"JET_DEV_TASK/JetDev/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()

	db := infrastructure.NewDatabase()
	// repo >> service >> controller >> routing >> server
	blogRepository := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepository)
	blogController := controller.NewBlogController(BlogService) // controller are being set up
	blogRoute := routes.NewBlogRoute(blogController, router)    // post routes are initialized
	blogRoute.Setup()                                           // post routes are being setup

	db.DB.AutoMigrate(&models.Article{}) // migrating Post model to datbase table
	router.Gin.Run(":8000")              //server started on 8000 port
}
