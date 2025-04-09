// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/account"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestAccountModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account.AccountModel)(nil)
	schema := account.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
