package dto

import "pos.com/app/errs"

type OrderRequest struct {
	Total      float64 `json:"total"`
	TotalItems int     `json:"totalItems"`
	Products   []OrderProductRequest
}

type OrderProductRequest struct {
	IdProduct int     `json:"idProduct"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderProduct struct {
	ProductId int     `json:"productId"`
	OrderId   int     `json:"orderId"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `json:"product" gorm:"foreignKey:Id;references:ProductId"`
}

type Order struct {
	Uuid          string         `json:"uuid"`
	Id            int            `json:"id"`
	Total         float64        `json:"total"`
	OrderProducts []OrderProduct `json:"order_products" gorm:"foreignKey:OrderId;references:Id"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
	DeletedAt     string         `json:"deleted_at"`
}

func (o OrderRequest) Validate() *errs.AppError {
	if o.Total <= 0 {
		return errs.NewValidationError("El total no puede ser menor o igual a cero")
	}
	if len(o.Products) == 0 {
		return errs.NewValidationError("Productos no validos")
	}
	return nil
}
