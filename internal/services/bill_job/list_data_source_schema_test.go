// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_job"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillJobsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*bill_job.BillJobsDataSourceModel)(nil)
	schema := bill_job.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
