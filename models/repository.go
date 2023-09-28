package models

import "time"

type Repository struct {
	ID              int
	Name            string
	FullName        string
	CloneURL        string
	LastUpdated     string
	LastUpdatedTime time.Time
}
