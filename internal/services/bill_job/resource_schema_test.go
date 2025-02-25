// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_job"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillJobModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bill_job.BillJobModel)(nil)
	schema := bill_job.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
