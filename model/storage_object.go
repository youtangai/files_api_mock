package model

const (
	// BLOB is type of file
	BLOB = "storage_blob"
	// TREE is type of directory
	TREE = "storage_tree"
)

type StorageObject struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
}

