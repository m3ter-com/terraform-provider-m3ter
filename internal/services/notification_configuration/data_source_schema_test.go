// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/notification_configuration"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestNotificationConfigurationDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*notification_configuration.NotificationConfigurationDataSourceModel)(nil)
  schema := notification_configuration.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
