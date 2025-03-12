// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/external_mapping"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestExternalMappingModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*external_mapping.ExternalMappingModel)(nil)
  schema := external_mapping.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
