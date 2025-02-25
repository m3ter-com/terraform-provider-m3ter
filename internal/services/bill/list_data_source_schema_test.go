// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bill.BillsDataSourceModel)(nil)
	schema := bill.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
