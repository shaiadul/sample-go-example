package product

type Product struct {
	ID          int     `json:"id"` // if i use id is make privete then it will give error , if need small case then use tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}
