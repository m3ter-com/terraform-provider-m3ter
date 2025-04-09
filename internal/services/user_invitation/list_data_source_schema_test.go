// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_invitation_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/user_invitation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestUserInvitationsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*user_invitation.UserInvitationsDataSourceModel)(nil)
	schema := user_invitation.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
