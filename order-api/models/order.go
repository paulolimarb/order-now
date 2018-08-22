package models

type OrderRequest struct {
	Order order `json:"order"`
}

type OrderResponse struct {
	OrderId      string `json:"order_id,omitempty"`
	DeliveryTime string `json:"delivery_time,omitempty"`
	Errors       errors `json:"errors,omitempty"`
}

type order struct {
	OrderReference string  `json:"order_reference"`
	Product        product `json:"product"`
	Buyer          buyer   `json:"buyer"`
}

type product struct {
	ProductName    string `json:"product_name"`
	Specifications string `json:"specifications"`
	PriceInCents   int64  `json:"price_in_cents"`
}

type buyer struct {
	BuyerName      string `json:"buyer_name"`
	DocumentNumber string `json:"document_number"`
	BuyerAddress   string `json:"buyer_address"`
	BuyerEmail     string `json:"buyer_email"`
}

type errors struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
