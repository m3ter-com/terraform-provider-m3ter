// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_config_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_config"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillConfigModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bill_config.BillConfigModel)(nil)
	schema := bill_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
