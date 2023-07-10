package dao

type Category struct {
	ID   int
	Name string
}

type Item struct {
	Id          int
	Name        string
	Description string
	CategoryID  int
	Category    Category `gorm:"foreignkey:CategoryID"`
	Price       float64
	Image       string
}
