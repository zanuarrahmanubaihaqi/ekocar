package model

type AddUserRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}
