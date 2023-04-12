package product

import (
	"sample-service/apperrors"
	"sample-service/domain"
	"sample-service/domain/category"
	"sample-service/domain/product"
)

// ProductDtoとProduct Entityの相互変換Adapter
// 2023/03/29
type ProductDtoConverter struct{}

// domain.EntityAdapterインターフェースのメソッド
// Product EntityからProductDtoを生成する
func (converter *ProductDtoConverter) Convert(entity any) (any, error) {
	source, ok := entity.(*product.Product)
	if !ok {
		return nil, apperrors.NewDomainError("指定されたEntityはProductではありません。")
	}
	// ProductDtoのインスタンスを生成する
	if source.Category() != nil { // 商品カテゴリが含まれている
		dto := NewProductDto(source.ProductId().Value(), source.ProductName().Value(), source.ProductPrice().Value(), source.Category().CategoryId().Value())
		return dto, nil
	} else { // 商品カテゴリが含まれていない
		dto := NewProductDto(source.ProductId().Value(), source.ProductName().Value(), source.ProductPrice().Value(), "")
		return dto, nil
	}
}

// domain.EntityAdapterインターフェースのメソッド
// ProductDtoからProduct Entityを再構築する
func (converter *ProductDtoConverter) Restore(model any) (any, error) {
	source, ok := model.(*ProductDto)
	if !ok {
		return nil, apperrors.NewDomainError("指定されたmodelはProductDtoではありません。")
	}
	var c *category.Category
	if source.CategoryId != "" { // 商品カテゴリIDが含まれている
		// Category Entityを生成する
		c, _ = category.BuildCategory(source.CategoryId, "")
	}
	// ProductDtoからProduct Entityを再構築する
	var p *product.Product
	var err error
	if source.Id == "" { // 商品IDが無い場合(商品の登録)
		p, err = product.NewProduct(source.Name, source.Price, c)
	} else {
		p, err = product.BuildProduct(source.Id, source.Name, source.Price, c)
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

// domain.EntitiesAdapterインターフェースのメソッド
// ProductのスライスからProductDtoのスライスを生成して返す
// 2023/03/29
func (converter *ProductDtoConverter) MultiConvert(entities any) (any, error) {
	products, ok := entities.([]product.Product)
	if !ok {
		return nil, apperrors.NewDomainError("指定されたentitiesは[]Productではありません。")
	}
	// ProductDtoを格納するスライスを生成する
	var dtos = make([]ProductDto, 0, len(products))
	// 引数productsからProductDtoのスライスを生成する
	for _, product := range products {
		categoryid := ""
		if product.Category() != nil { // 商品カテゴリが含まれている
			categoryid = product.Category().CategoryId().Value()
		}
		dto := NewProductDto(product.ProductId().Value(), product.ProductName().Value(), product.ProductPrice().Value(), categoryid)
		dtos = append(dtos, *dto)
	}
	return dtos, nil
}

// domain.EntitiesAdapterインターフェースのメソッド
// ProductDyoのスライスからProductのスライスを生成して返す
// 2023/03/29
func (converter *ProductDtoConverter) MultiRestore(models any) (any, error) {
	return nil, nil
}

// コンストラクタ
func NewProductDtoConverter() domain.EntitiesConverter {
	return &ProductDtoConverter{}
}
