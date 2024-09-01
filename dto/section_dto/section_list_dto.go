package section_dto

import (
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/user_dto"
	"time"
)

type SectionListingDTO struct {
	ID          uint                        `json:"id"`
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	CreatedAt   time.Time                   `json:"created_at"`
	User        user_dto.UserBasicDTO       `json:"user"`
	Project     project_dto.ProjectBasicDTO `json:"project"`
}
