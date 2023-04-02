package dto

// 商品を扱うDTO
// 2023/03/28
type ProductDto struct {
	Id    string `josn:"id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

// 　コンストラクタ
func NewProductDto(id string, name string, price uint32) *ProductDto {
	return &ProductDto{Id: id, Name: name, Price: price}
}
