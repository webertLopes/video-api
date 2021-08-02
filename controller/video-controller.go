package controller

import (
	"interface/entity"
	"interface/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	GetAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	videoService service.VideoService
}

func New(videoService service.VideoService) VideoController {
	return &videoController{
		videoService: videoService,
	}
}

func (controller *videoController) GetAll() []entity.Video {
	return controller.videoService.FindAll()
}

func (controller *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	controller.videoService.Create(video)
	return video
}
