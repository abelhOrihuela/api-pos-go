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
	Id            int     `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid          string  `gorm:"unique;not null;type:varchar(100);default:null" db:"uuid"`
	TotalItems    int     `gorm:"not null;type:int;default:null" db:"total_items"`
	Total         float64 `gorm:"not null;type:double precision;default:null" db:"total"`
	OrderProducts []OrderProduct
}

type OrderProduct struct {
	gorm.Model
	Id        int     `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid      string  `gorm:"unique;not null;type:varchar(100)" db:"uuid"`
	OrderId   string  `gorm:"not null;type:int;default:null" db:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderId;references:Id"`
	Total     float64 `gorm:"not null;type:double precision;default:null" db:"total"`
	Price     float64 `gorm:"not null;type:double precision;default:0" db:"price"`
	Quantity  float64 `gorm:"not null;type:double precision;default:0" db:"quantity"`
	ProductId int     `gorm:"not null;type:int;default:null" db:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId;references:Id"`
}

func CreateOrder(req dto.OrderRequest) (*Order, *errs.AppError) {

	errValidation := req.Validate()

	if errValidation != nil {
		return nil, errValidation
	}

	var products []OrderProduct
	var productsUpdated []Product

	for _, p := range req.Products {

		var product Product

		err := db.Database.Where(&Product{Id: p.IdProduct}).First(&product).Error

		if err != nil {
			return nil, errs.NewNotFoundError("Product not found")
		}

		products = append(products, OrderProduct{
			Total:     p.Price * float64(p.Quantity),
			Quantity:  p.Quantity,
			ProductId: product.Id,
			Price:     p.Price,
		})

		product.CurrentExistence -= float64(p.Quantity)

		productsUpdated = append(productsUpdated, product)
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

	for _, p := range productsUpdated {
		db.Database.Save(&p)
	}

	return &p, nil

}

func GetOrder(uuid string) (*Order, *errs.AppError) {
	var order Order

	err := db.Database.Where("uuid = ?", uuid).Preload("OrderProducts").First(&order).Error

	if err != nil {
		return nil, errs.NewDefaultError(err.Error())
	}

	return &order, nil
}

func GetAllOrders(req *http.Request) paginate.Page {

	model := db.Database.Order("created_at DESC").Preload("OrderProducts").Preload("OrderProducts.Product").Model(&Order{})
	pg := paginate.New()

	page := pg.With(model).Request(req).Response(&[]dto.Order{})

	return page
}

func (order *Order) BeforeCreate(*gorm.DB) error {
	order.Uuid = uuid.NewString()
	return nil
}

func (orderProduct *OrderProduct) BeforeCreate(*gorm.DB) error {
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
		Uuid:          o.Uuid,
		Id:            o.Id,
		Total:         o.Total,
		OrderProducts: products,
	}
}
