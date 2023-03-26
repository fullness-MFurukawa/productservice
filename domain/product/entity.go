package product

import (
	"fmt"
	"reflect"
	"sample-service/domain"
	"sample-service/domain/category"

	"github.com/google/uuid"
)

// 商品を表すEntity
// 2023/03/25
type Product struct {
	productId    *ProductId         // 商品番号
	productName  *ProductName       // 商品名
	productPrice *ProductPrice      // 単価
	category     *category.Category // カテゴリ
}

// コンストラクタ
func NewProduct(name string, price uint32, category *category.Category) (*Product, error) {
	uid, u_err := uuid.NewRandom()
	if u_err != nil {
		return nil, domain.NewDomainError(u_err.Error())
	}
	id, err := NewProductId(uid.String())
	if err != nil {
		return nil, err
	}
	product_name, err := NewProductName(name)
	if err != nil {
		return nil, err
	}
	product_price, err := NewProductPrice(price)
	if err != nil {
		return nil, err
	}
	return &Product{productId: id, productName: product_name, productPrice: product_price, category: category}, nil
}

// 商品Entityの再構築
func BuildProduct(product_id string, product_name string, product_price uint32, category *category.Category) (*Product, error) {
	id, err := NewProductId(product_id)
	if err != nil {
		return nil, err
	}
	name, err := NewProductName(product_name)
	if err != nil {
		return nil, err
	}
	price, err := NewProductPrice(product_price)
	if err != nil {
		return nil, err
	}
	return &Product{productId: id, productName: name, productPrice: price, category: category}, nil
}

// ゲッター
func (entity *Product) ProductId() *ProductId {
	return entity.productId
}
func (entity *Product) ProductName() *ProductName {
	return entity.productName
}
func (entity *Product) ProductPrice() *ProductPrice {
	return entity.productPrice
}
func (entity *Product) Category() *category.Category {
	return entity.category
}

// インスタンスの内容
func (entity Product) String() string {
	var category_str = ""
	if entity.category != nil {
		category_str = entity.category.String()
	}
	return fmt.Sprintf("Product[ProductId:%s,ProductName:%s,ProductPrice:%d, %s]",
		entity.productId.value, entity.productName.value, entity.productPrice.value, category_str)
}

// 等価性検証メソッド
func (entity *Product) Equals(obj interface{}) (bool, error) {
	object, ok := obj.(*Product)
	if !ok {
		return false, domain.NewDomainError("Productの等価検証で、異なる型が指定されました。")
	}
	// DeepEqualsで値を比較する
	result := reflect.DeepEqual(entity.productId, object.productId)
	return result, nil
}
