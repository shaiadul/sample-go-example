package repo

type Product struct {
	ID          int     `json:"id"` // if i use id is make privete then it will give error , if need small case then use tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(productID int) error
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor fucntion
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProduct(repo)
	return repo
}

func (pr *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(pr.productList) + 1
	pr.productList = append(pr.productList, &p)
	return &p, nil
}

func (pr *productRepo) Get(productID int) (*Product, error) {
	for _, product := range pr.productList {
		if product.ID == productID {
			return product, nil
		}
	}

	return nil, nil
}

func (pr *productRepo) List() ([]*Product, error) {
	return pr.productList, nil
}

func (pr *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range pr.productList {
		if product.ID == p.ID {
			pr.productList[idx] = &product
		}
	}

	return &product, nil
}

func (pr *productRepo) Delete(productID int) error {
	for idx, product := range pr.productList {
		if product.ID == productID {
			pr.productList = append(pr.productList[:idx], pr.productList[idx+1:]...)
			break
		}
	}
	return nil
}

func generateInitialProduct(r *productRepo) {
	pro01 := &Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is product 1 description",
		Price:       100.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro02 := &Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is product 2 description",
		Price:       200.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro03 := &Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is product 3 description",
		Price:       300.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro04 := &Product{
		ID:          4,
		Title:       "Product 4",
		Description: "This is product 4 description",
		Price:       400.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	pro05 := &Product{
		ID:          5,
		Title:       "Product 5",
		Description: "This is product 5 description",
		Price:       500.00,
		ImgUrl:      "https://via.placeholder.com/150",
	}

	r.productList = append(r.productList, pro01, pro02, pro03, pro04, pro05)
}
