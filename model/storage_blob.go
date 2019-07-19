package model

type StorageBlob struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
	Data string `json:"data"`
}
