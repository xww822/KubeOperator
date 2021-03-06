package repository

import (
	"github.com/KubeOperator/KubeOperator/pkg/db"
	"github.com/KubeOperator/KubeOperator/pkg/model"
)

type SystemRegistry interface {
	Get(arch string) (model.SystemRegistry, error)
	List() ([]model.SystemRegistry, error)
	Save(registry *model.SystemRegistry) error
}

type systemRegistry struct {
}

func (s systemRegistry) Get(arch string) (model.SystemRegistry, error) {
	var registry model.SystemRegistry
	if err := db.DB.Where(&model.SystemRegistry{Architecture: arch}).First(&registry).Error; err != nil {
		return registry, err
	}
	return registry, nil
}

func (s systemRegistry) List() ([]model.SystemRegistry, error) {
	var registry []model.SystemRegistry
	err := db.DB.Model(&model.SystemRegistry{}).Find(&registry).Error
	return registry, err
}

func (s systemRegistry) Save(registry *model.SystemRegistry) error {
	if db.DB.NewRecord(registry) {
		return db.DB.Create(&registry).Error
	} else {
		return db.DB.Model(&registry).Updates(&registry).Error
	}
}
