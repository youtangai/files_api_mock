package service

import (
	"encoding/base64"
	"fmt"
	"github.com/youtangai/files_api_mock/config"
	"github.com/youtangai/files_api_mock/model"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	wd string
)

func init() {
	log.Println("file service: info: start initializing")

	targetDir := config.GetWD()

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
	ReadFile(path string) (model.StorageBlob, error)
	ReadDir(path string) (model.StorageTree, error)
	DeleteNode(path string) error
	IsDir(path string) (bool, error)
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

// CreateDir is Create new dir in the specific locale.
func (srv FileService) CreateDir(path string) error {
	targetPath := filepath.Join(wd, path)
	err := os.Mkdir(targetPath, 0755)
	if err != nil {
		return fmt.Errorf("file service: err: failed to create dir: path: %s, err: %s¥n", targetPath, err)
	}
	return nil
}

func (srv FileService) ReadFile(path string) (model.StorageBlob, error) {
	targetPath := filepath.Join(wd, path)
	file, err := os.Open(targetPath)
	if err != nil {
		return model.StorageBlob{}, fmt.Errorf("file service: err: failed to open file: path: %s, err: %s¥n", targetPath, err)
	}
	defer file.Close()

	// len, cap is 1MB
	data := make([]byte, 1024*1024)
	content := ""
	for {
		_, err := file.Read(data)
		// EOFがerrに入ったら，すべて内容を読み込んだということで終了
		if err == io.EOF {
			break
		}
		if err != nil {
			return model.StorageBlob{}, fmt.Errorf("file service: err: failed to read content: path: %s, err: %s¥n", targetPath, err)
		}
		content += string(data)
	}

	//base64でエンコード
	encoded := base64.StdEncoding.EncodeToString([]byte(content))

	return model.StorageBlob{
		Kind: "Blob",
		Path: path,
		Data: encoded,
	}, nil
}

func (srv FileService) ReadDir(path string) (model.StorageTree, error) {
	targetPath := filepath.Join(wd, path)
	nodes, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return model.StorageTree{}, fmt.Errorf("file service: err: failed to read dir: path: %s, err: %s¥n", targetPath, err)
	}

	// 長さがわかっているので，明示的に指定
	items := make([]model.StorageObject, len(nodes))

	for index, node := range nodes {
		// ディレクトリだったら Kind:Tree を追加
		if node.IsDir() {
			items[index] = model.StorageObject{
				Kind: "Tree",
				Path: filepath.Join(path, node.Name()),
			}
			continue
		}
		// ディレクトリでなければ Kind:Blob を追加
		items[index] = model.StorageObject{
			Kind: "Blob",
			Path: filepath.Join(path,node.Name()),
		}
	}

	return model.StorageTree{
		Kind: "Tree",
		Path: path,
		Items: items,
	}, nil
}

func (srv FileService) DeleteNode(path string) error {
	targetPath := filepath.Join(wd, path)
	err := os.RemoveAll(targetPath)
	if err != nil {
		return fmt.Errorf("file service: err: failed to delete dir: path: %s, err: %s¥n", targetPath, err)
	}
	return nil
}

func (srv FileService) IsDir(path string) (bool, error) {
	targetPath := filepath.Join(wd, path)
	file, err := os.Open(targetPath)
	if err != nil {
		return false, fmt.Errorf("file service: err: failed to open path: %s, err: %s¥n", targetPath, err)
	}

	info, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("file service: err: failed to get file info: %s, err: %s¥n", targetPath, err)
	}

	if info.IsDir() {
		return true, nil
	}

	return false, nil
}