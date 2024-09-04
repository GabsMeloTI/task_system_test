package service

import (
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"gorm.io/gorm"
)

type SectionService interface {
	GetSections() ([]section_dto.SectionListingDTO, error)
	GetSectionByID(id uint) (section_dto.SectionListingDTO, error)
	CreateSection(section models.Section) error
	UpdateSection(id uint, section models.Section) error
	DeleteSection(id uint) error
}

type sectionService struct {
	db *gorm.DB
}

func NewSectionService(db *gorm.DB) SectionService {
	return &sectionService{db: db}
}

func (s *sectionService) GetSections() ([]section_dto.SectionListingDTO, error) {
	var sections []models.Section
	var sectionsDTO []section_dto.SectionListingDTO

	err := s.db.Preload("Project").Preload("User").Find(&sections).Error
	if err != nil {
		return nil, err
	}

	for _, section := range sections {
		sectionDTO := section_dto.SectionListingDTO{
			ID:          section.ID,
			Title:       section.Title,
			Description: section.Description,
			CreatedAt:   section.CreatedAt,
			User: user_dto.UserBasicDTO{
				ID:    section.User.ID,
				Name:  section.User.Name,
				Email: section.User.Email,
			},
			Project: project_dto.ProjectBasicDTO{
				ID:     section.Project.ID,
				Title:  section.Project.Title,
				Status: section.Project.Status,
			},
		}
		sectionsDTO = append(sectionsDTO, sectionDTO)
	}

	return sectionsDTO, nil
}

func (s *sectionService) GetSectionByID(id uint) (section_dto.SectionListingDTO, error) {
	var section models.Section

	err := s.db.Preload("Project").Preload("User").First(&section, id).Error
	if err != nil {
		return section_dto.SectionListingDTO{}, err
	}

	sectionDTO := section_dto.SectionListingDTO{
		ID:          section.ID,
		Title:       section.Title,
		Description: section.Description,
		CreatedAt:   section.CreatedAt,
		User: user_dto.UserBasicDTO{
			ID:    section.User.ID,
			Name:  section.User.Name,
			Email: section.User.Email,
		},
		Project: project_dto.ProjectBasicDTO{
			ID:     section.Project.ID,
			Title:  section.Project.Title,
			Status: section.Project.Status,
		},
	}

	return sectionDTO, nil
}

func (s *sectionService) CreateSection(section models.Section) error {
	return s.db.Create(&section).Error
}

func (s *sectionService) UpdateSection(id uint, section models.Section) error {
	var existingSection models.Section

	err := s.db.First(&existingSection, id).Error
	if err != nil {
		return err
	}

	return s.db.Model(&existingSection).Updates(section).Error
}

func (s *sectionService) DeleteSection(id uint) error {
	return s.db.Delete(&models.Section{}, id).Error
}
