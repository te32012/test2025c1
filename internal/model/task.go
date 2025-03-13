package model

import "time"

type Task struct {
	Id int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	Created_at time.Time `json:"created_at,omitzero"`
	Updated_at time.Time `json:"updated_at,omitzero"`
}