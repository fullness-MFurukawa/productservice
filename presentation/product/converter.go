package product

import (
	"sample-service/apperrors"
	"sample-service/domain"
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
	dto := NewProductDto(source.ProductId().Value(), source.ProductName().Value(), source.ProductPrice().Value())
	return dto, nil
}

// domain.EntityAdapterインターフェースのメソッド
// ProductDtoからProduct Entityを再構築する
func (converter *ProductDtoConverter) Restore(model any) (any, error) {
	source, ok := model.(*ProductDto)
	if !ok {
		return nil, apperrors.NewDomainError("指定されたmodelはProductDtoではありません。")
	}
	product, err := product.BuildProduct(source.Id, source.Name, source.Price, nil)
	if err != nil {
		return nil, err
	}
	return *product, nil
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
		dto := NewProductDto(product.ProductId().Value(), product.ProductName().Value(), product.ProductPrice().Value())
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
