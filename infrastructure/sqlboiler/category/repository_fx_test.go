package category

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestCategoryRepository(t *testing.T) {
	app := fxtest.New(t,
		fx.Options(CategoryModule),
	)
	defer app.RequireStop()
}
