package model

type Menu struct {
	Model
	Name   string `json:"name"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
