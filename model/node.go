package model

const (
	// BLOB is type of file
	BLOB = "blob"
	// TREE is type of directory
	TREE = "Tree"
)

type Node struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

