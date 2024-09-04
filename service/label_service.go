package service

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/models"
	"gorm.io/gorm"
)

type LabelService struct {
	DB *gorm.DB
}

func NewLabelService(db *gorm.DB) *LabelService {
	return &LabelService{DB: db}
}

func (s *LabelService) GetAllLabels() ([]label_dto.LabelListingDTO, error) {
	var labels []models.Label
	var labelsDTO []label_dto.LabelListingDTO

	if err := db.DB.Find(&labels).Error; err != nil {
		return nil, err
	}

	for _, label := range labels {
		labelDTO := label_dto.LabelListingDTO{
			ID:    label.ID,
			Name:  label.Name,
			Color: label.Color,
		}
		labelsDTO = append(labelsDTO, labelDTO)
	}

	return labelsDTO, nil
}

func (s *LabelService) GetLabelByID(id uint) (label_dto.LabelListingDTO, error) {
	var label models.Label

	if err := db.DB.First(&label, id).Error; err != nil {
		return label_dto.LabelListingDTO{}, err
	}

	labelDTO := label_dto.LabelListingDTO{
		ID:    label.ID,
		Name:  label.Name,
		Color: label.Color,
	}

	return labelDTO, nil
}

func (s *LabelService) CreateLabel(label models.Label) error {
	return db.DB.Create(&label).Error
}

func (s *LabelService) UpdateLabel(id uint, label models.Label) error {
	var existingLabel models.Label

	if err := db.DB.First(&existingLabel, id).Error; err != nil {
		return err
	}

	existingLabel.Name = label.Name
	existingLabel.Color = label.Color

	return db.DB.Save(&existingLabel).Error
}

func (s *LabelService) DeleteLabel(id uint) error {
	return db.DB.Delete(&models.Label{}, id).Error
}
