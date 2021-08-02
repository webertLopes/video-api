package service

import (
	"interface/entity"
	"interface/repository"
)

type VideoService interface {
	FindAll() []entity.Video
	Create(entity.Video) entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: videoRepository,
	}
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.GetAll()
}

func (service *videoService) Create(video entity.Video) entity.Video {
	go service.videoRepository.Save(video)
	return video
}
