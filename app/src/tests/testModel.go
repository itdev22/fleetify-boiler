package tests

type TestModel struct {
	ID              int    `gorm:"column:id" json:"id"`
	CustomerName    string `gorm:"column:customer_name" json:"customerName"`
	CustomerPhone   string `gorm:"column:customer_phone" json:"customerPhone"`
	CustomerEmail   string `gorm:"column:customer_email" json:"customerEmail"`
	CustomerAddress string `gorm:"column:customer_address" json:"customerAddress"`
}
