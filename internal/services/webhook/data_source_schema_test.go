// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/webhook"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestWebhookDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*webhook.WebhookDataSourceModel)(nil)
	schema := webhook.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
