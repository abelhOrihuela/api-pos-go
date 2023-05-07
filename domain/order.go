package domain

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"pos.com/app/db"
	"pos.com/app/dto"
	"pos.com/app/errs"
)

type Order struct {
	gorm.Model
	Id            int            `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid          string         `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	TotalItems    int            `gorm:"not null;type:int;default:null" db:"total_items"`
	Total         float64        `gorm:"not null;type:double precision;default:null" db:"total"`
	OrderProducts []OrderProduct `gorm:"foreignKey:OrderId;references:Id"`
}

type OrderProduct struct {
	gorm.Model
	Id        int     `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid      string  `gorm:"unique;not null;type:varchar(100)" db:"uuid"`
	OrderId   string  `gorm:"not null;type:int;default:null" db:"order_id"`
	Total     float64 `gorm:"not null;type:double precision;default:null" db:"total"`
	Quantity  int16   `gorm:"not null;type:int;default:null" db:"quantity"`
	ProductId int     `gorm:"not null;type:int;default:null" db:"product_id"`
}

func CreateOrder(req dto.OrderRequest) (*Order, *errs.AppError) {

	errValidation := req.Validate()

	if errValidation != nil {
		return nil, errValidation
	}

	var products []OrderProduct

	for _, p := range req.Products {

		var product Product

		err := db.Database.Where(&Product{Id: p.IdProduct}).First(&product).Error

		if err != nil {
			return nil, errs.NewNotFoundError("Product not found")
		}

		products = append(products, OrderProduct{
			Total:     p.Price,
			Quantity:  p.Quantity,
			ProductId: product.Id,
		})
	}

	p := Order{
		TotalItems:    req.TotalItems,
		Total:         req.Total,
		OrderProducts: products,
	}

	err := db.Database.Create(&p).Error

	if err != nil {
		return nil, errs.NewUnexpectedDatabaseError("Unexpected error during the creation of order" + err.Error())
	}
	return &p, nil

}

func GetAllOrders(req *http.Request) paginate.Page {

	model := db.Database.Preload("OrderProducts").Preload("OrderProducts.Product").Model(&Order{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.Order{})

	return page
}

func (order *Order) BeforeCreate(*gorm.DB) error {

	order.Uuid = uuid.NewString()
	return nil
}

func (orderProduct *OrderProduct) BeforeSave(*gorm.DB) error {

	orderProduct.Uuid = uuid.NewString()
	return nil
}
func (o Order) ToDto() dto.Order {
	var products []dto.OrderProduct

	for _, p := range o.OrderProducts {
		products = append(products, dto.OrderProduct{
			OrderId:   o.Id,
			Quantity:  p.Quantity,
			ProductId: p.ProductId,
		})
	}
	return dto.Order{
		Id:            o.Id,
		Total:         o.Total,
		OrderProducts: products,
	}
}
