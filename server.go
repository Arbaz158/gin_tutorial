package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_tutorial/controller"
	"github.com/gin_tutorial/middlewares"
	"github.com/gin_tutorial/service"
)

var (
	VideoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(VideoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})
	server.Run(":8080")
}
