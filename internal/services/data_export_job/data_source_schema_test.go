// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_job_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/data_export_job"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestDataExportJobDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*data_export_job.DataExportJobDataSourceModel)(nil)
  schema := data_export_job.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
