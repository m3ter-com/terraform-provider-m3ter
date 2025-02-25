// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/data_export_schedule"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestDataExportScheduleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*data_export_schedule.DataExportScheduleModel)(nil)
	schema := data_export_schedule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
