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

func TestFileService_DeleteNode(t *testing.T) {
	srv := NewFileService()
	err := srv.DeleteNode("a.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileService_ReadDir(t *testing.T) {
	srv := NewFileService()
	tree, err := srv.ReadDir("dirA/dirB")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v¥n",tree)
}

func TestFileService_ReadFile(t *testing.T) {
	srv :=NewFileService()
	blob, err := srv.ReadFile("a.txt")
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := base64.StdEncoding.DecodeString(blob.Data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s¥n", decoded)
}