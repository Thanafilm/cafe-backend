package model

type Food struct {
	FoodID    string
	Name      string
	UnitPrice float32
	Category  string
}

type Order struct {
	TableID    string
	OrderID    string
	FoodID     string
	Date       string
	TotalPrice float32
}

type Payment struct {
	PaymentID string
	TableID   string
	OrderID   string
	Date      string
	Amount    float32
}
