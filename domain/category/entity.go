package category

import (
	"fmt"
	"reflect"
	"sample-service/domain"

	"github.com/google/uuid"
)

// 商品カテゴリを表すEntity
type Category struct {
	categoryId   *CategoryId   // カテゴリ番号
	categoryName *CategoryName // カテゴリ名
}

// コンストラクタ
func NewCategory(category_name string) (*Category, error) {
	uid, u_err := uuid.NewRandom()
	if u_err != nil {
		return nil, domain.NewDomainError(u_err.Error())
	}
	id, err := NewCategoryId(uid.String())
	if err != nil {
		return nil, err
	}
	name, err := NewCategoryName(category_name)
	if err != nil {
		return nil, err
	}
	return &Category{categoryId: id, categoryName: name}, nil
}

// Entityの再構築
func BuildCategory(categtory_id string, category_name string) (*Category, error) {
	id, err := NewCategoryId(categtory_id)
	if err != nil {
		return nil, err
	}
	name, err := NewCategoryName(category_name)
	if err != nil {
		return nil, err
	}
	return &Category{categoryId: id, categoryName: name}, nil
}

// ゲッター
func (entity *Category) CategoryId() CategoryId {
	return *entity.categoryId
}
func (entity *Category) CategoryName() CategoryName {
	return *entity.categoryName
}

// インスタンスの内容
func (entity Category) String() string {
	return fmt.Sprintf("Category[CategoryId:%s,CategoryName:%s]",
		entity.categoryId.value, entity.categoryName.value)
}

// 等価性検証メソッド
func (entity *Category) Equals(obj interface{}) (bool, error) {
	object, ok := obj.(*Category)
	if !ok {
		return false, domain.NewDomainError("Categoryの等価検証で、異なる型が指定されました。")
	}
	// DeepEqualsで値を比較する
	result := reflect.DeepEqual(entity.categoryId, object.categoryId)
	return result, nil
}
