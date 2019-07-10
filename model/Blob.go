package model

type Blob struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Data string `json:"data"`
}
