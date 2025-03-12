// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/scheduled_event_configuration"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestScheduledEventConfigurationsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*scheduled_event_configuration.ScheduledEventConfigurationsDataSourceModel)(nil)
  schema := scheduled_event_configuration.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
