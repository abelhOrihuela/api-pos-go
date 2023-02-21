package dto

type ProductResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Barcode     string `json:"barcode"`
	Price       int64  `json:"price"`
}
