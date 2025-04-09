// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_credit_line_item_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_credit_line_item"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillCreditLineItemDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bill_credit_line_item.BillCreditLineItemDataSourceModel)(nil)
	schema := bill_credit_line_item.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
