package service

import (
	_ "awesomeProject/db"
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"gorm.io/gorm"
)

type ProjectService interface {
	GetProjects() ([]project_dto.ProjectListingDTO, error)
	GetProjectByID(id uint) (project_dto.ProjectListingDTO, error)
	CreateProject(project models.Project) error
	UpdateProject(id uint, project models.Project) error
	DeleteProject(id uint) error
}

type projectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) ProjectService {
	return &projectService{db: db}
}

func (s *projectService) GetProjects() ([]project_dto.ProjectListingDTO, error) {
	var projects []models.Project
	var projectsDTO []project_dto.ProjectListingDTO

	err := s.db.Preload("User").Find(&projects).Error
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		projectDTO := project_dto.ProjectListingDTO{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			Status:      project.Status,
			CreatedAt:   project.CreatedAt,
			User: user_dto.UserBasicDTO{
				ID:   project.User.ID,
				Name: project.User.Name,
			},
		}
		projectsDTO = append(projectsDTO, projectDTO)
	}

	return projectsDTO, nil
}

func (s *projectService) GetProjectByID(id uint) (project_dto.ProjectListingDTO, error) {
	var project models.Project

	err := s.db.Preload("User").Preload("Sections").First(&project, id).Error
	if err != nil {
		return project_dto.ProjectListingDTO{}, err
	}

	projectDTO := project_dto.ProjectListingDTO{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		User: user_dto.UserBasicDTO{
			ID:   project.User.ID,
			Name: project.User.Name,
		},
	}

	return projectDTO, nil
}

func (s *projectService) CreateProject(project models.Project) error {
	return s.db.Create(&project).Error
}

func (s *projectService) UpdateProject(id uint, project models.Project) error {
	var existingProject models.Project

	err := s.db.First(&existingProject, id).Error
	if err != nil {
		return err
	}

	return s.db.Model(&existingProject).Updates(project).Error
}

func (s *projectService) DeleteProject(id uint) error {
	return s.db.Delete(&models.Project{}, id).Error
}
