package domain

type Product struct {
	ID        int    `json:"ID"`
	Name      string `json:"Name"`
	Image_url string `json:"Image"`
	Price     string `json:"Price"`
	Create_By int    `json:"Create_UserID"`
}

type ProductRepository interface {
	GetByID(id int) (*Product, error)
	GetByPage(page_size int, page_number int) ([]Product, error)
	Create(product *Product) error
	UpdateByID(id int, product *Product) error
	DeleteByID(id int) error
}

type ProductInteractor interface {
	GetByID(id int) (*Product, error)
	GetByPage(page_size int, page_number int) ([]Product, error)
	Create(product *Product) error
	UpdateByID(id int, product *Product) error
	DeleteByID(id int) error
}
