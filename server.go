package main

import (
	"interface/controller"
	"interface/entity"
	"interface/infra"
	"interface/repository"
	"interface/service"

	"github.com/gin-gonic/gin"
)

var connection = "root:root@/digi-docs"
var database = infra.ConnectToDb(connection)

var (
	videoRepository repository.VideoRepository = repository.New(database)
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	database.AutoMigrate(&entity.Video{})

	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.GetAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	defer database.Close()

	server.Run(":5000")
}
