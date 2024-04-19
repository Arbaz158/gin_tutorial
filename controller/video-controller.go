package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin_tutorial/entity"
	"github.com/gin_tutorial/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "video page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
