// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestAggregationsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*aggregation.AggregationsDataSourceModel)(nil)
	schema := aggregation.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
