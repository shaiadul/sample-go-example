package repo

type Product struct {
	ID          int     `json:"id"` // if i use id is make privete then it will give error , if need small case then use tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create()
	Get()
	List()
	Update()
	Delete()
}

type productRepo struct {
	productList []Product
}

// constructor or constructor fucntion
func NewProductRepo() ProductRepo {
	return productRepo{}
}

// Create method to create a new product
func (pr productRepo) Create() {}
func (pr productRepo) Get()    {}
func (pr productRepo) List()   {}
func (pr productRepo) Update() {}
func (pr productRepo) Delete() {}
