package model

type Tree struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Items []Node `json:"items"`
}
