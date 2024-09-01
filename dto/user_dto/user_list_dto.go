package user_dto

import "time"

type UserListingDTO struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Photo        string    `json:"photo"`
	CreatedAt    time.Time `json:"created_at"`
	ProjectCount int       `json:"project_count"`
}
