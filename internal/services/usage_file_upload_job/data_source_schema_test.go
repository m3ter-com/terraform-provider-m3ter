// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package usage_file_upload_job_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/usage_file_upload_job"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestUsageFileUploadJobDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*usage_file_upload_job.UsageFileUploadJobDataSourceModel)(nil)
  schema := usage_file_upload_job.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
