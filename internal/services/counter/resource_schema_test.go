// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCounterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*counter.CounterModel)(nil)
	schema := counter.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
