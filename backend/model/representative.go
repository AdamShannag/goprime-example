package model

type Representative struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}
