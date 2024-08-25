package model

type Country struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}
