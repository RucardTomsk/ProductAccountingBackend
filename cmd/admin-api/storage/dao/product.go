package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"productAccounting-v1/internal/domain/entity"
)

type ProductStorage struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) *ProductStorage {
	return &ProductStorage{
		db: db,
	}
}

func (s *ProductStorage) CreateProduct(product *entity.Product) error {
	return s.db.Create(&product).Error
}

func (s *ProductStorage) RetrieveProduct(id *uuid.UUID) (*entity.Product, error) {
	var product entity.Product
	err := s.db.First(&product, id).Error
	return &product, err
}

func (s *ProductStorage) UpdateProduct(product *entity.Product) error {
	tx := s.db.Updates(product)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *ProductStorage) GetProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := s.db.Model(&entity.Product{}).
		Preload("Assembly").
		Find(&products).Error

	return products, err
}

func (s *ProductStorage) CreateAssembly(assembly *entity.Assembly) error {
	return s.db.Create(&assembly).Error
}

func (s *ProductStorage) UpdateAssembly(assembly *entity.Assembly) error {
	tx := s.db.Updates(assembly)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *ProductStorage) RetrieveAssembly(id *uuid.UUID) (*entity.Assembly, error) {
	var assembly entity.Assembly
	err := s.db.First(&assembly, id).Error
	return &assembly, err
}

func (s *ProductStorage) GetAssemblyComponents(id *uuid.UUID) ([]entity.Component, error) {
	var assembly entity.Assembly
	err := s.db.Model(&entity.Assembly{}).
		Preload("Components").
		First(&assembly, id).Error

	return assembly.Components, err
}

func (s *ProductStorage) DeleteProduct(id *uuid.UUID) error {
	tx := s.db.Delete(&entity.Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
