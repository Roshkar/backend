package pkg

type Order struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Products []Product `json:"products"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
}

type Product struct {
	ID       string
	Name     string
	Category string
	Quantity int
	Price    float64
}

type OrderedProduct struct {
	ID              string
	ProductId       string
	ProductQuantity int
	OrderId         string
}

type ExampleOrderRequest struct {
	Name     string `default:"Ivan Ivanov"`
	Address  string `default:"Sofia Mladost 2"`
	Phone    string `default:"0888888888"`
	Products []struct {
		Id       string `default:"bc264186-9c2e-4533-6ba5-705c160303c1"`
		Quantity int    `default:"2"`
	}
}

type ExampleProductRequest struct {
	Name     string  `default:"Men Red Shirt"`
	Category string  `default:"Men Shirts"`
	Quantity int     `default:"1000"`
	Price    float64 `default:"19.99"`
}
