package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/youtangai/files_api_mock/controller"
	"net/http"
)

func main() {
	router, err := initRouter()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(router.Run())
}

func initRouter() (*gin.Engine, error) {
	fileCtrl := controller.NewFileController()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	router.GET("/files/:path", fileCtrl.GetNodes)
	router.POST("/files/:path", fileCtrl.CreateNode)
	router.DELETE("/files/*path", fileCtrl.DeleteNode)

	return router, nil
}