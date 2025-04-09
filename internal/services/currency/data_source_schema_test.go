// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/currency"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCurrencyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*currency.CurrencyDataSourceModel)(nil)
	schema := currency.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
