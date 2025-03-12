// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/custom_field"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCustomFieldModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*custom_field.CustomFieldModel)(nil)
  schema := custom_field.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
