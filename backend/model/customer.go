package model

type Customer struct {
	ID             int            `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	Country        Country        `json:"country"`
	Company        string         `json:"company,omitempty"`
	Date           string         `json:"date,omitempty"`
	Status         string         `json:"status,omitempty"`
	Verified       bool           `json:"verified,omitempty"`
	Activity       int            `json:"activity,omitempty"`
	Balance        float64        `json:"balance,omitempty"`
	Representative Representative `json:"representative"`
}
