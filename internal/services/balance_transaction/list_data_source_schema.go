// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance_transaction

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BalanceTransactionsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "balance_id": schema.StringAttribute{
        Required: true,
      },
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "transaction_type_id": schema.StringAttribute{
        Optional: true,
      },
      "max_items": schema.Int64Attribute{
        Description: "Max items to fetch, default: 1000",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.AtLeast(0),
        },
      },
      "items": schema.ListNestedAttribute{
        Description: "The items returned by the data source",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[BalanceTransactionsItemsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The UUID of the entity.",
              Computed: true,
            },
            "version": schema.Int64Attribute{
              Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
              Computed: true,
            },
            "amount": schema.Float64Attribute{
              Description: "The financial value of the transaction, as recorded in the balance.",
              Computed: true,
            },
            "applied_date": schema.StringAttribute{
              Description: "The date *(in ISO 8601 format)* when the balance transaction was applied, i.e., when the balance was affected.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "created_by": schema.StringAttribute{
              Description: "The unique identifier (UUID) for the user who created the balance transaction.",
              Computed: true,
            },
            "currency_paid": schema.StringAttribute{
              Description: "The currency code such as USD, GBP, EUR of the payment, if it differs from the balance currency.",
              Computed: true,
            },
            "description": schema.StringAttribute{
              Description: "A brief description explaining the purpose or context of the transaction.",
              Computed: true,
            },
            "dt_created": schema.StringAttribute{
              Description: "The date and time *(in ISO 8601 format)* when the balance transaction was first created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The date and time *(in ISO 8601 format)* when the balance transaction was last modified.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "entity_id": schema.StringAttribute{
              Description: "The unique identifier (UUID) for the entity associated with the Transaction, as specified by the `entityType`.",
              Computed: true,
            },
            "entity_type": schema.StringAttribute{
              Description: "The type of entity associated with the Transaction - identifies who or what was responsible for the Transaction being added to the Balance - such as a **User**, a **Service User**, or a **Bill**.\nAvailable values: \"BILL\", \"COMMITMENT\", \"USER\", \"SERVICE_USER\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "BILL",
                "COMMITMENT",
                "USER",
                "SERVICE_USER",
              ),
              },
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The unique identifier (UUID) for the user who last modified the balance transaction.",
              Computed: true,
            },
            "paid": schema.Float64Attribute{
              Description: "The actual payment amount if the payment currency differs from the Balance currency.",
              Computed: true,
            },
            "transaction_date": schema.StringAttribute{
              Description: "The date *(in ISO 8601 format)* when the transaction was recorded in the system.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "transaction_type_id": schema.StringAttribute{
              Description: "The unique identifier (UUID) for the Transaction type. This is obtained from the list of created Transaction Types within the Organization Configuration.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *BalanceTransactionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *BalanceTransactionsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
