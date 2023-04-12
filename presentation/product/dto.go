package product

// 商品を扱うDTO
// 2023/03/28
type ProductDto struct {
	Id         string `josn:"id"`
	Name       string `json:"name"`
	Price      uint32 `json:"price"`
	CategoryId string `json:"category_id"`
}

// 　コンストラクタ
func NewProductDto(id string, name string, price uint32, categoryid string) *ProductDto {
	return &ProductDto{Id: id, Name: name, Price: price, CategoryId: categoryid}
}
