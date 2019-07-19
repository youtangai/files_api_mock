package model

type StorageTree struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
	Items []StorageObject `json:"items"`
}
