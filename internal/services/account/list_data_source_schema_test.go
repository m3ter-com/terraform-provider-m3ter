// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/account"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestAccountsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account.AccountsDataSourceModel)(nil)
	schema := account.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
