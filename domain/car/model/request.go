package model

type AddCarRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}

type UpdateCarRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}
