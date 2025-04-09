// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/compound_aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCompoundAggregationsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compound_aggregation.CompoundAggregationsDataSourceModel)(nil)
	schema := compound_aggregation.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
