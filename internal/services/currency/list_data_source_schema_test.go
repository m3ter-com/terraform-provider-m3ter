// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/currency"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCurrenciesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*currency.CurrenciesDataSourceModel)(nil)
	schema := currency.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
