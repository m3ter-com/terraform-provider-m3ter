// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_pricing_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter_pricing"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCounterPricingsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*counter_pricing.CounterPricingsDataSourceModel)(nil)
	schema := counter_pricing.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
