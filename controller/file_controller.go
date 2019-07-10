package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/youtangai/files_api_mock/model"
	"github.com/youtangai/files_api_mock/service"
	"net/http"
	"fmt"
)

type IFileController interface {
	GetNodes(c *gin.Context)
	CreateNode(c *gin.Context)
	DeleteNode(c *gin.Context)
}

type FileController struct {
	Srv service.IFileService
}

func NewFileController(srv service.IFileService) IFileController {
	return FileController{
		Srv: srv,
	}
}

func (ctrl FileController) GetNodes(c *gin.Context) {
	path := c.Param("path")
	fmt.Println(path)

	items := []model.Node{
		{
			Kind: model.BLOB,
			Name: "file1",
		},
		{
			Kind: model.BLOB,
			Name: "file2",
		},
	}
	result := model.Tree{
		Kind: model.TREE,
		Name: "/",
		Items: items,
	}
	c.JSON(http.StatusOK, result)
}

func (ctrl FileController) CreateNode(c *gin.Context) {
	path := c.Param("path")
	fmt.Println(path)
	c.Status(http.StatusCreated)
}

func (ctrl FileController) DeleteNode(c *gin.Context) {
	path := c.Param("path")
	fmt.Println(path)
	c.Status(http.StatusNoContent)
}