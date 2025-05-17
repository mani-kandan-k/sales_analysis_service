package refreshdata

type OrderDetailsStruct struct {
	Id            int     `gorm:"column:id"`
	OrderID       string  `gorm:"column:order_id"`
	CustomerId    string  `gorm:"-"`
	ProductId     string  `gorm:"-"`
	CustomerKeyId int     `gorm:"column:customer_id"`
	ProductKeyId  int     `gorm:"column:product_id"`
	Region        string  `gorm:"column:region"`
	Category      string  `gorm:"column:category"`
	DateOfSale    string  `gorm:"column:date_of_sale"`
	QuantitySold  int     `gorm:"column:quantity_sold"`
	ShippingCost  float64 `gorm:"column:shipping_cost"`
	PaymentMethod string  `gorm:"column:payment_method"`
	CreatedDate   string  `gorm:"column:CreatedDate"`
	UpdatedDate   string  `gorm:"column:UpdatedDate"`
}

type CustomerDetailsStruct struct {
	Id              int    `gorm:"column:id"`
	CustomerID      string `gorm:"column:customer_id"`
	CustomerName    string `gorm:"column:customer_name"`
	CustomerEmail   string `gorm:"column:customer_email"`
	CustomerAddress string `gorm:"column:customer_address"`
	CreatedDate     string `gorm:"column:CreatedDate"`
	UpdatedDate     string `gorm:"column:UpdatedDate"`
}

type ProductDetailsStruct struct {
	Id          int     `gorm:"column:id"`
	ProductID   string  `gorm:"column:product_id"`
	ProductName string  `gorm:"column:product_name"`
	UnitPrice   float64 `gorm:"column:unit_price"`
	Discount    float64 `gorm:"column:discount"`
	CreatedDate string  `gorm:"column:CreatedDate"`
	UpdatedDate string  `gorm:"column:UpdatedDate"`
}
