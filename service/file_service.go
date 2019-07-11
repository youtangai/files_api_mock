package service

import (
	"encoding/base64"
	"fmt"
	"github.com/youtangai/files_api_mock/model"
	"log"
	"os"
	"path/filepath"
)

const (
	targetDir = "../data"
)

var (
	wd string
)

func init() {
	log.Println("file service: info: start initializing")

	err := os.Chdir(targetDir)
	if err != nil {
		log.Fatal("file service: err: failed change directory:", err)
	}

	wd, err = os.Getwd()
	if err != nil {
		log.Fatal("file service: err: failed get current working directory:", err)
	}
	log.Println("file service: info: working directory is", wd)
}

type IFileService interface {
	CreateFile(path, content string) error
	CreateDir(path string) error
	ReadFile(path string) (model.Blob, error)
	ReadDir(path string) (model.Tree, error)
	DeleteFile(path string) error
	DeleteDir(path string) error
}

type FileService struct {}

func NewFileService() IFileService {
	return FileService{}
}

// CreateFile is Create new file in the specific locale. content is encoded with base64.
func (srv FileService) CreateFile(path, content string) error {
	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return fmt.Errorf("file service: err: failed to decode content: %s¥n", err)
	}

	targetPath := filepath.Join(wd, path)
	file, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("file service: err: failed to create file: path: %s, err: %s¥n", targetPath, err)
	}

	_, err = file.Write(decoded)
	if err != nil {
		return fmt.Errorf("file service: err: failed to write content: err: %s¥n", err)
	}
	return nil
}

func (srv FileService) CreateDir(path string) error {
	fmt.Println(path)
	return nil
}

func (srv FileService) ReadFile(path string) (model.Blob, error) {
	fmt.Println(path)
	return model.Blob{}, nil
}

func (srv FileService) ReadDir(path string) (model.Tree, error) {
	fmt.Println(path)
	return model.Tree{}, nil
}

func (srv FileService) DeleteFile(path string) error {
	fmt.Println(path)
	return nil
}

func (srv FileService) DeleteDir(path string) error {
	fmt.Println(path)
	return nil
}