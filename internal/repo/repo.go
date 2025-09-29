package repo

import (
	"gorm.io/gorm"
	"marketplace/internal/models"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepo) GetAll(products *[]models.Product) error {
	return r.DB.Find(products).Error
}

func (r *ProductRepo) GetById(product *models.Product, id string) error {
	return r.DB.First(product, id).Error
}

func (r *ProductRepo) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepo) Delete(id string) error {
	return r.DB.Delete(&models.Product{}, id).Error
}
