package dto

import "pos.com/app/errs"

type OrderRequest struct {
	Total      float64 `json:"total"`
	TotalItems int     `json:"totalItems"`
	Products   []OrderProductRequest
}

type OrderProductRequest struct {
	IdProduct uint    `json:"idProduct"`
	Quantity  int16   `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderProduct struct {
	ProductId uint    `json:"productId"`
	OrderId   uint    `json:"orderId"`
	Quantity  int16   `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `json:"product" gorm:"foreignKey:Id;references:ProductId"`
}

type Order struct {
	Id            uint           `json:"id"`
	Total         float64        `json:"total"`
	OrderProducts []OrderProduct `json:"order_products" gorm:"foreignKey:OrderId;references:Id"`
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
