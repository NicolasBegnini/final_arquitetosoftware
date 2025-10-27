package repository

import (
	config "github.com/NicolasBegnini/final_arquitetosoftware.git/database"
)

type ProductRepository struct{}

func (r ProductRepository) Create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func (r ProductRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	result := config.DB.First(&product, id)
	return product, result.Error
}

func (r ProductRepository) FindByName(name string) ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Where("name LIKE ?", "%"+name+"%").Find(&products)
	return products, result.Error
}

func (r ProductRepository) Update(product *models.Product) error {
	return config.DB.Save(product).Error
}

func (r ProductRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}

func (r ProductRepository) Count() (int64, error) {
	var count int64
	result := config.DB.Model(&models.Product{}).Count(&count)
	return count, result.Error
}
