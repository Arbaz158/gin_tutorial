package service

import (
	"github.com/gin_tutorial/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type Service struct {
	videos []entity.Video
}

func New() VideoService {
	return &Service{}
}

func (service *Service) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *Service) FindAll() []entity.Video {
	return service.videos
}
