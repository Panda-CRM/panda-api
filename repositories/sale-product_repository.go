package repositories

import (
	"github.com/Panda-CRM/panda-api/database"
	"github.com/Panda-CRM/panda-api/helpers"
	"github.com/Panda-CRM/panda-api/models"
	"log"
	"net/url"
)

type SaleProductRepository interface {
	List(q url.Values) models.SaleProductList
	Get(id string) models.SaleProduct
	GetByDocument(document string) models.SaleProduct
	Delete(id string) error
	Create(p *models.SaleProduct) error
	Update(p *models.SaleProduct) error
	CountRows() int
}

type saleProductRepository struct{}

func NewSaleProductRepository() *saleProductRepository {
	return new(saleProductRepository)
}

func (repository saleProductRepository) List(q url.Values) models.SaleProductList {
	db := database.GetInstance()
	pageParams := helpers.MakePagination(repository.CountRows(), q.Get("page"), q.Get("per_page"))
	var saleProductsList models.SaleProductList
	saleProductsList.Pages = pageParams

	db.Limit(pageParams.ItemPerPage).
		Offset(pageParams.StartIndex).
		Find(&saleProductsList.SaleProducts)

	return saleProductsList
}

func (repository saleProductRepository) Get(id string) models.SaleProduct {
	db := database.GetInstance()
	var saleProduct models.SaleProduct
	db.Where("uuid = ?", id).First(&saleProduct)
	return saleProduct
}

func (repository saleProductRepository) GetByDocument(document string) models.SaleProduct {
	db := database.GetInstance()
	var product models.SaleProduct
	db.Where("document = ?", document).First(&product)
	return product
}

func (repository saleProductRepository) Delete(id string) error {
	db := database.GetInstance()

	err := db.Where("uuid = ?", id).Delete(&models.SaleProduct{}).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository saleProductRepository) Create(s *models.SaleProduct) error {
	db := database.GetInstance()

	err := db.Create(&s).Error
	if err != nil {
		log.Print(err.Error())
	}
	return err
}

func (repository saleProductRepository) Update(s *models.SaleProduct) error {
	db := database.GetInstance()

	err := db.Model(&s).
		Omit("uuid").
		Updates(&s).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository saleProductRepository) CountRows() int {
	db := database.GetInstance()
	var count int
	db.Model(&models.SaleProduct{}).Count(&count)
	return count
}
