package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"` // if i use id is make privete then it will give error , if need small case then use tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func init() {
	pro01 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is product 1 description",
		Price:       100.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro02 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is product 2 description",
		Price:       200.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro03 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is product 3 description",
		Price:       300.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro04 := Product{
		ID:          4,
		Title:       "Product 4",
		Description: "This is product 4 description",
		Price:       400.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro05 := Product{
		ID:          5,
		Title:       "Product 5",
		Description: "This is product 5 description",
		Price:       500.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	ProductList = append(ProductList, pro01, pro02, pro03, pro04, pro05)
}
