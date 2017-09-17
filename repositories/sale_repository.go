package repositories

import (
	"github.com/Panda-CRM/panda-api/database"
	"github.com/Panda-CRM/panda-api/helpers"
	"github.com/Panda-CRM/panda-api/models"
	"log"
	"net/url"
)

type SaleRepository interface {
	List(q url.Values) models.SaleList
	Get(id string) models.Sale
	Delete(id string) error
	Create(p *models.Sale) error
	Update(p *models.Sale) error
	CountRows() int
}

type saleRepository struct{}

func NewSaleRepository() *saleRepository {
	return new(saleRepository)
}

func (repository saleRepository) List(q url.Values) models.SaleList {
	db := database.GetInstance()
	pageParams := helpers.MakePagination(repository.CountRows(), q.Get("page"), q.Get("per_page"))
	var salesList models.SaleList
	salesList.Pages = pageParams

	/*db.Raw(`
	SELECT
		s.uuid, s.code, s.registered_at,
		s.sale_date, s.seller_uuid, s.buyer_uuid,
		b.name, se.name,
		SUM(sp.product_value) as total_value
	FROM
		sales AS s
	INNER JOIN
		sale_products AS sp
		ON (sp.sale_uuid = s.uuid)
	INNER JOIN
		people AS b
		ON (b.uuid = s.buyer_uuid)
	INNER JOIN
		people AS se
		ON (se.uuid = s.seller_uuid)
	GROUP BY
		s.uuid,
		s.code,
		b.name,
		se.name
	ORDER BY
		s.sale_date DESC;`).Scan(&salesList.Sales)*/

	db.
		Debug().Table("sales AS s").
		Select(`
			s.uuid, s.code, s.registered_at, s.sale_date,
			s.seller_uuid, s.buyer_uuid,
			SUM(sp.product_value) AS total_value`).
		Joins("INNER JOIN sale_products AS sp ON sp.sale_uuid = s.uuid").
		Group("s.uuid, s.code").
		Preload("Buyer").
		Preload("Seller").
		Limit(pageParams.ItemPerPage).
		Offset(pageParams.StartIndex).
		Order("sale_date desc").
		Find(&salesList.Sales)

	return salesList
}

func (repository saleRepository) Get(id string) models.Sale {
	db := database.GetInstance()
	var sale models.Sale
	db.Where("uuid = ?", id).First(&sale)
	return sale
}

func (repository saleRepository) Delete(id string) error {
	db := database.GetInstance()

	err := db.Where("uuid = ?", id).Delete(&models.Sale{}).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository saleRepository) Create(s *models.Sale) error {
	db := database.GetInstance()

	err := db.Set("gorm:save_associations", false).
		Create(&s).Error
	if err != nil {
		log.Print(err.Error())
	}
	return err
}

func (repository saleRepository) Update(s *models.Sale) error {
	db := database.GetInstance()

	err := db.Model(&s).
		Omit("uuid").
		Updates(&s).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository saleRepository) CountRows() int {
	db := database.GetInstance()
	var count int
	db.Model(&models.Sale{}).Count(&count)
	return count
}
