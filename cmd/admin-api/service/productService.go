package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/storage/dao"
	"productAccounting-v1/internal/domain/base"
	"productAccounting-v1/internal/domain/entity"
)

type ProductService struct {
	storage          *dao.ProductStorage
	componentStorage *dao.ComponentStorage
}

func NewProductService(
	storage *dao.ProductStorage,
	componentStorage *dao.ComponentStorage) *ProductService {
	return &ProductService{
		storage:          storage,
		componentStorage: componentStorage,
	}
}

func (s *ProductService) CreateProduct(request *model.CreateProductRequest) (*uuid.UUID, *base.ServiceError) {
	product := &entity.Product{
		Name:        request.Name,
		Description: request.Description,
	}

	if err := s.storage.CreateProduct(product); err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return &product.ID, nil
}

func (s *ProductService) UpdateProduct(productID *uuid.UUID, request *model.UpdateProductRequest) *base.ServiceError {
	product, err := s.storage.RetrieveProduct(productID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}

		return base.NewPostgresReadError(err)
	}

	product.Name = request.Name
	product.Description = request.Description

	if err := s.storage.UpdateProduct(product); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

func (s *ProductService) AddAssembly(productID *uuid.UUID, request *model.CreateAssemblyRequest) *base.ServiceError {
	product, err := s.storage.RetrieveProduct(productID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	assembly := entity.Assembly{
		Name:   request.Name,
		Weight: request.Weight,
	}

	product.Assembly = append(product.Assembly, assembly)

	if err := s.storage.UpdateProduct(product); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

func (s *ProductService) AddComponentToAssembly(assemblyID, componentID *uuid.UUID) *base.ServiceError {
	assembly, err := s.storage.RetrieveAssembly(assemblyID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	component, err := s.componentStorage.RetrieveComponent(componentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return base.NewNotFoundError(err)
		}
		return base.NewPostgresReadError(err)
	}

	assembly.Components = append(assembly.Components, *component)

	if err := s.storage.UpdateAssembly(assembly); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

func (s *ProductService) GetProduct() ([]model.ProductObject, *base.ServiceError) {
	products, err := s.storage.GetProducts()
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	result := make([]model.ProductObject, len(products))

	for _, product := range products {
		assembles := make([]model.AssemblyObject, len(product.Assembly))
		for _, assembly := range product.Assembly {
			assembles = append(assembles, model.AssemblyObject{
				ID:     assembly.ID,
				Name:   assembly.Name,
				Weight: assembly.Weight,
			})
		}
		result = append(result, model.ProductObject{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Assembles:   assembles,
		})
	}

	return result, nil
}

func (s *ProductService) GetAssemblyComponent(id *uuid.UUID) ([]model.ComponentObject, *base.ServiceError) {
	components, err := s.storage.GetAssemblyComponents(id)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	result := make([]model.ComponentObject, len(components))

	for _, component := range components {
		result = append(result, model.ComponentObject{
			ID:     component.ID,
			Name:   component.Name,
			Weight: component.Weight,
			Price:  component.Weight,
		})
	}

	return result, nil
}
