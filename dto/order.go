package dto

type Order struct {
	Id         uint    `gorm:"primaryKey;autoIncrement" db:"id"`
	Uuid       string  `db:"uuid"`
	Amount     float64 `db:"name"`
	TotalItems int     `db:"price"`
}

type OrderResponse struct {
	Id         uint    `json:"id"`
	Uuid       string  `json:"uuid"`
	Amount     float64 `json:"description"`
	Barcode    string  `json:"barcode"`
	TotalItems int     `json:"price"`
}

type OrderRequest struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price"`
	Products    []ProductResponse
}
