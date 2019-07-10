package service

import (
	"fmt"
	"github.com/youtangai/files_api_mock/model"
)

type IFileService interface {
	CreateFile(name, content string) error
	CreateDir(name string) error
	ReadFile(name string) (model.Blob, error)
	ReadDir(name string) (model.Tree, error)
	DeleteFile(name string) error
	DeleteDir(name string) error
}

type FileService struct {}

func NewFileService() IFileService {
	return FileService{}
}

func (srv FileService) CreateFile(name, content string) error {
	fmt.Println(name, content)
	return nil
}

func (srv FileService) CreateDir(name string) error {
	fmt.Println(name)
	return nil
}

func (srv FileService) ReadFile(name string) (model.Blob, error) {
	fmt.Println(name)
	return model.Blob{}, nil
}

func (srv FileService) ReadDir(name string) (model.Tree, error) {
	fmt.Println(name)
	return model.Tree{}, nil
}

func (srv FileService) DeleteFile(name string) error {
	fmt.Println(name)
	return nil
}

func (srv FileService) DeleteDir(name string) error {
	fmt.Println(name)
	return nil
}