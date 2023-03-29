package domain

// Entityと他のモデルを変換するConverterインターフェース
// 2023-03-23
type EntityConverter interface {
	// 任意のEntityから任意のModelへ変換する
	Convert(entity any) (any, error)
	// 任意のModeleから任意のEntityへ変換する
	Restore(model any) (any, error)
}

// 複数のEntityと他の複数のモデルを変換するConverterインターフェース
// 2023-03-29
type EntitiesConverter interface {
	EntityConverter // EntityConverterインターフェースを埋め込む
	// 任意の複数のEntityから任意の複数Modelへ変換する
	MultiConvert(entities any) (any, error)
	// 任意の複数のModeleから任意の複数のEntityへ変換する
	MultiRestore(models any) (any, error)
}
