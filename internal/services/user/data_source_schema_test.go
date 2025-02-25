// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/user"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestUserDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user.UserDataSourceModel)(nil)
	schema := user.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
