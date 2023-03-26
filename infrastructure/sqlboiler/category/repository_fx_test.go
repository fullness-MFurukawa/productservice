package category

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/infrastructure/sqlboiler/tests"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func testSelectAll(rep category.CategoryRepository) {
	ctx, transaction := tests.TestDBInit()
	results, _ := rep.FindAll(ctx, transaction)
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestCategoryRepository(t *testing.T) {
	app := fxtest.New(t,
		fx.Options(CategoryModule),
		fx.Invoke(testSelectAll),
	)
	defer app.RequireStart().RequireStop()
	assert.True(t, true)
}
