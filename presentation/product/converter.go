package product

import (
	"sample-service/domain"
	"sample-service/domain/product"
)

// ProductDtoとProduct Entityの相互変換Adapter
// 2023/03/29
type ProductDtoConverter struct{}

// domain.EntityAdapterインターフェースのメソッド
// Product EntityからProductDtoを生成する
func (adapter *ProductDtoConverter) Convert(entity interface{}) (interface{}, error) {
	source, ok := entity.(*product.Product)
	if !ok {
		return nil, domain.NewDomainError("指定されたEntityはProductではありません。")
	}
	// ProductDtoのインスタンスを生成する
	dto := NewProductDto(source.ProductId().Value(), source.ProductName().Value(), source.ProductPrice().Value())
	return dto, nil
}

// domain.EntityAdapterインターフェースのメソッド
// ProductDtoからProduct Entityを再構築する
func (adapter *ProductDtoConverter) Restore(model interface{}) (interface{}, error) {
	source, ok := model.(*ProductDto)
	if !ok {
		return nil, domain.NewDomainError("指定されたmodelはProductDtoではありません。")
	}
	product, err := product.BuildProduct(source.Id, source.Name, source.Price, nil)
	if err != nil {
		return nil, err
	}
	return *product, nil
}

// ProductDtoAdapterのメソッド
// ProductのスライスからProductDtoのスライスを生成して返す
// 2023/03/29
func (adapter *ProductDtoConverter) Converts(products []*product.Product) []*ProductDto {
	// ProductDtoを格納するスライスを生成する
	var dtos = make([]*ProductDto, 0, len(products))
	// 引数productsからProductDtoのスライスを生成する
	for _, product := range products {
		dto := NewProductDto(product.ProductId().Value(), product.ProductName().Value(), product.ProductPrice().Value())
		dtos = append(dtos, dto)
	}
	return dtos
}

// コンストラクタ
func NewProductDtoConverter() domain.EntityConverter {
	return &ProductDtoConverter{}
}
