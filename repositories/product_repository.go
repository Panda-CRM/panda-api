package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/models"
	"log"
	"net/url"
	"strconv"
)

type ProductRepository interface {
	List(q url.Values) models.ProductList
	Get(id string) models.Product
	Delete(id string) error
	Create(p *models.Product) error
	Update(p *models.Product) error
	CountRows() int
}

type productRepository struct{}

func NewProductRepository() *productRepository {
	return new(productRepository)
}

func (repository productRepository) List(q url.Values) models.ProductList {
	db := database.GetInstance()

	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)

	var productsList models.ProductList
	productsList.Meta.Pagination = pagination

	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("registered_at desc").
		Find(&productsList.Products)

	return productsList
}

func (repository productRepository) Get(id string) models.Product {
	db := database.GetInstance()
	var product models.Product
	db.Where("uuid = ?", id).First(&product)
	return product
}

func (repository productRepository) Delete(id string) error {
	db := database.GetInstance()

	err := db.Where("uuid = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository productRepository) Create(p *models.Product) error {
	db := database.GetInstance()

	err := db.Create(&p).Error
	if err != nil {
		log.Print(err.Error())
	}
	return err
}

func (repository productRepository) Update(p *models.Product) error {
	db := database.GetInstance()

	err := db.Model(&p).
		Omit("uuid").
		Updates(&p).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository productRepository) CountRows() int {
	db := database.GetInstance()
	var count int
	db.Model(&models.Product{}).Count(&count)
	return count
}
