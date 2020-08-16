package model

import "time"

type Service struct {
	Id           int
	Name         string
	ProjectId    int
	RelativePath string
	CreatedAt    time.Time
	ModifiedAt   time.Time
	ModifiedBy   string
}
