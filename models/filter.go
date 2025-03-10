package models

type Filter struct {
	Name   string         `json:"name"`
	Values *[]interface{} `json:"values,omitempty"`
	From   interface{}    `json:"from,omitempty"`
	To     interface{}    `json:"to,omitempty"`
}
