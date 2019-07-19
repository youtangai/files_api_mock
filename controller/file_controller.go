package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/youtangai/files_api_mock/model"
	"github.com/youtangai/files_api_mock/service"
	"net/http"
)

const (
	BLOB = "blob"
	TREE = "tree"
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

	isDir, err := ctrl.Srv.IsDir(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if isDir {
		tree, err := ctrl.Srv.ReadDir(path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK,tree)
		return
	}

	blob, err := ctrl.Srv.ReadFile(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, blob)
}

func (ctrl FileController) CreateNode(c *gin.Context) {
	path := c.Param("path")
	fmt.Println(path)

	var json model.Blob
	var err error
	err = c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	switch json.Kind {
	case BLOB:
		err = ctrl.Srv.CreateFile(path, json.Data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		break
	case TREE:
		err = ctrl.Srv.CreateDir(path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		break
	}
	c.Status(http.StatusCreated)
}

func (ctrl FileController) DeleteNode(c *gin.Context) {
	path := c.Param("path")
	fmt.Println(path)

	err := ctrl.Srv.DeleteNode(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}