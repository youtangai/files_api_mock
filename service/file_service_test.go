package service

import (
	"encoding/base64"
	"testing"
)

func TestCreateFile(t *testing.T) {
	srv := NewFileService()
	content := []byte("project winter")
	encoded := base64.StdEncoding.EncodeToString(content)

	err := srv.CreateFile("a.txt", encoded)
	if err != nil {
		t.Fatal(err)
	}
}