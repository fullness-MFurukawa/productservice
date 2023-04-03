package category

// 商品カテゴリを扱うDTO
// 2023/03/28
type CategoryDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// コンストラクタ
func NewCategoryDto(id string, name string) *CategoryDto {
	return &CategoryDto{Id: id, Name: name}
}
