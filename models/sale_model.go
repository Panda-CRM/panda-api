package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"time"
)

type Sale struct {
	UUID         string        `json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	Code         int           `json:"code,omitempty" sql:"auto_increment; primary_key; unique"`
	RegisteredAt *time.Time    `json:"registered_at,omitempty" sql:"type:timestamp without time zone; default:NOW()"`
	SaleDate     time.Time     `json:"sale_date" sql:"type:timestamp without time zone; default:NOW()"`
	SellerUUID   string        `json:"-" sql:"type:uuid; default:uuid_nil()"`
	Seller       Person        `json:"seller"`
	BuyerUUID    string        `json:"-" sql:"type:uuid; default:uuid_nil()"`
	Buyer        Person        `json:"buyer"`
	Products     []SaleProduct `json:"products"`
	TotalValue   float32       `json:"total_value" gorm:"-"`
}

type SaleProduct struct {
	UUID                        string     `json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	SaleUUID                    string     `json:"-" sql:"type:uuid; not null"`
	Document                    string     `json:"document" sql:"type:varchar(50)"`
	ProductValue                float32    `json:"product_value" sql:"type:numeric"`
	TaxValue                    float32    `json:"tax_value" sql:"type:numeric"`
	AgencyDiscountValue         float32    `json:"agency_discount_value" sql:"type:numeric"`
	SupplierDiscountValue       float32    `json:"supplier_discount_value" sql:"type:numeric"`
	CommissionValue             float32    `json:"commission_value" sql:"type:numeric"`
	ValueIntermediaryCommission float32    `json:"value_intermediary_commission" sql:"type:numeric"`
	ScriptDescription           string     `json:"script_description" sql:"type:varchar(50)"`
	ProductCode                 string     `json:"-" sql:"type:varchar(20)"`
	ProductDescription          string     `json:"product_description" sql:"type:varchar(50)"`
	DescriptionProductType      string     `json:"-" sql:"type:varchar(50)"`
	CommissionPercentage        float32    `json:"commission_percentage" sql:"type:numeric"`
	ProfitValue                 float32    `json:"profit_value" sql:"type:numeric"`
	TotalValue                  float32    `json:"total_value" sql:"type:numeric"`
	DateCancellation            *time.Time `json:"date_cancellation,omitempty" sql:"type:timestamp without time zone; default:null"`
	DateShipment                time.Time  `json:"date_shipment,omitempty" sql:"type:timestamp without time zone; default:null"`
	ReturnDate                  time.Time  `json:"return_date,omitempty" sql:"type:timestamp without time zone; default:null"`
}

type SaleList struct {
	Sales []Sale       `json:"sales"`
	Pages   helpers.PageParams `json:"pages"`
}

type SaleProductList struct {
	SaleProducts []SaleProduct `json:"sale_products"`
	Pages   helpers.PageParams `json:"pages"`
}

func (s Sale) IsEmpty() bool {
	return s.UUID == ""
}

func (s Sale) Validate() []string {
	var errs []string
	if _, err := govalidator.ValidateStruct(s); err != nil {
		errsV := err.(govalidator.Errors).Errors()
		for _, element := range errsV {
			errs = append(errs, element.Error())
		}
	}
	return errs
}

func (s *Sale) BeforeCreate(scope *gorm.Scope) error {
	dateTime := time.Now()
	scope.SetColumn("code", nil)
	scope.SetColumn("registered_at", dateTime)
	scope.SetColumn("uuid", nil)
	return nil
}

func (sp SaleProduct) IsEmpty() bool {
	return sp == SaleProduct{}
}
