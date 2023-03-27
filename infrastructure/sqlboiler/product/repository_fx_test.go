package product

import (
	"fmt"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/tests"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func testSelectAll(rep product.ProductRepositiry) {
	ctx, transaction := tests.TestDBInit()
	products, err := rep.FindAll(ctx, transaction)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}
func TestProductRepository(t *testing.T) {
	app := fxtest.New(t,
		fx.Options(Module),
		fx.Invoke(testSelectAll),
	)
	defer app.RequireStart().RequireStop()
	assert.True(t, true)
}
