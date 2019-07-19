package service

import (
	"encoding/base64"
	"testing"
)

func TestFileService_CreateFile(t *testing.T) {
	srv := NewFileService()
	content := []byte("project winter")
	encoded := base64.StdEncoding.EncodeToString(content)

	err := srv.CreateFile("a.txt", encoded)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileService_CreateDir(t *testing.T) {
	srv := NewFileService()
	err := srv.CreateDir("dir1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileService_DeleteFile(t *testing.T) {
	srv := NewFileService()
	err := srv.DeleteFile("a.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileService_DeleteDir(t *testing.T) {
	srv := NewFileService()
	err := srv.DeleteDir("dir1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileService_ReadDir(t *testing.T) {
	srv := NewFileService()
	tree, err := srv.ReadDir("")
	if err != nil {
		t.Fatal(err)
	}

	if len(tree.Items) != 3 {
		t.Fatal(err)
	}

	t.Logf("%v¥n",tree)
}