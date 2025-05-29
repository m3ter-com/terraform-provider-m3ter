// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_job_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/statement_statement_job"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestStatementStatementJobModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*statement_statement_job.StatementStatementJobModel)(nil)
	schema := statement_statement_job.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
