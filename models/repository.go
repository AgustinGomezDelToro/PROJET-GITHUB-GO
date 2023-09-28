package models

import "time"

type Repository struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	CloneURL        string    `json:"clone_url"`
	LastUpdated     string    `json:"updated_at"`
	LastUpdatedTime time.Time `json:"-"`
}
