// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "org_id": schema.StringAttribute{
        Optional: true,
      },
      "codes": schema.ListAttribute{
        Description: "List of Account Codes to retrieve. \nThese are unique short codes for each Account.",
        Optional: true,
        ElementType: types.StringType,
      },
      "ids": schema.ListAttribute{
        Description: "List of Account IDs to retrieve.",
        Optional: true,
        ElementType: types.StringType,
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
        CustomType: customfield.NewNestedObjectListType[AccountsItemsDataSourceModel](ctx),
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
            "address": schema.SingleNestedAttribute{
              Description: "Contact address.",
              Computed: true,
              CustomType: customfield.NewNestedObjectType[AccountsAddressDataSourceModel](ctx),
              Attributes: map[string]schema.Attribute{
                "address_line1": schema.StringAttribute{
                  Computed: true,
                },
                "address_line2": schema.StringAttribute{
                  Computed: true,
                },
                "address_line3": schema.StringAttribute{
                  Computed: true,
                },
                "address_line4": schema.StringAttribute{
                  Computed: true,
                },
                "country": schema.StringAttribute{
                  Computed: true,
                },
                "locality": schema.StringAttribute{
                  Computed: true,
                },
                "post_code": schema.StringAttribute{
                  Computed: true,
                },
                "region": schema.StringAttribute{
                  Computed: true,
                },
              },
            },
            "auto_generate_statement_mode": schema.StringAttribute{
              Description: "Specify whether to auto-generate statements once Bills are approved or locked.\n\n- **None**. Statements will not be auto-generated.\n- **JSON**. Statements are auto-generated in JSON format.\n- **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.\nAvailable values: \"NONE\", \"JSON\", \"JSON_AND_CSV\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "NONE",
                "JSON",
                "JSON_AND_CSV",
              ),
              },
            },
            "bill_epoch": schema.StringAttribute{
              Description: "Defines first bill date for Account Bills. For example, if the Plan attached to the Account is set for monthly billing frequency and you set the first bill date to be January 1st, Bills are created every month starting on that date.\n\nOptional attribute - if not defined, then first bill date is determined by the Epoch settings at Organizational level.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "code": schema.StringAttribute{
              Description: "Code of the Account.\nThis is a unique short code used for the Account.",
              Computed: true,
            },
            "config_data": schema.MapAttribute{
              Description: "Configuration data for the Account",
              Computed: true,
              CustomType: customfield.NewMapType[jsontypes.Normalized](ctx),
              ElementType: jsontypes.NormalizedType{

              },
            },
            "created_by": schema.StringAttribute{
              Description: "The ID of the user who created the account.",
              Computed: true,
            },
            "credit_application_order": schema.ListAttribute{
              Description: "The order in which any Prepayment or Balance amounts on the Account are to be drawn-down against for billing. Four options:\n- `\"PREPAYMENT\",\"BALANCE\"`. Draw-down against Prepayment credit before Balance credit.\n- `\"BALANCE\",\"PREPAYMENT\"`. Draw-down against Balance credit before Prepayment credit.\n- `\"PREPAYMENT\"`. Only draw-down against Prepayment credit.\n- `\"BALANCE\"`. Only draw-down against Balance credit.",
              Computed: true,
              Validators: []validator.List{
              listvalidator.ValueStringsAre(
                stringvalidator.OneOfCaseInsensitive("PREPAYMENT", "BALANCE"),
              ),
              },
              CustomType: customfield.NewListType[types.String](ctx),
              ElementType: types.StringType,
            },
            "currency": schema.StringAttribute{
              Description: "Account level billing currency, such as USD or GBP. Optional attribute:\n- If you define an Account currency, this will be used for bills.\n- If you do not define a currency, the billing currency defined at Organizational will be used.\n\n**Note:** If you've attached a Plan to the Account that uses a different currency to the billing currency, then you must add the relevant currency conversion rate at Organization level to ensure the billing process can convert line items calculated using the Plan currency into the selected billing currency. If you don't add these conversion rates, then bills will fail for the Account.",
              Computed: true,
            },
            "custom_fields": schema.MapAttribute{
              Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
              Computed: true,
              CustomType: customfield.NewMapType[types.Dynamic](ctx),
              ElementType: types.DynamicType,
            },
            "days_before_bill_due": schema.Int64Attribute{
              Description: "The number of days after the Bill generation date shown on Bills as the due date.",
              Computed: true,
            },
            "dt_created": schema.StringAttribute{
              Description: "The DateTime when the Account was created *(in ISO 8601 format)*.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The DateTime when the Account was last modified *(in ISO 8601 format)*.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "email_address": schema.StringAttribute{
              Description: "Contact email for the Account.",
              Computed: true,
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The ID of the user who last modified the Account.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "Name of the Account.",
              Computed: true,
            },
            "parent_account_id": schema.StringAttribute{
              Description: "Parent Account ID, or null if this account does not have a parent.",
              Computed: true,
            },
            "purchase_order_number": schema.StringAttribute{
              Description: "Purchase Order Number of the Account.\n\nOptional attribute - allows you to set a purchase order number that comes through into invoicing. For example, your financial systems might require this as a reference for clearing payments.",
              Computed: true,
            },
            "statement_definition_id": schema.StringAttribute{
              Description: "The UUID of the statement definition used when Bill statements are generated for the Account. If no statement definition is specified for the Account, the statement definition specified at Organizational level is used.\n\nBill statements can be used as informative backing sheets to invoices. Based on the usage breakdown defined in the statement definition, generated statements give a breakdown of usage charges on Account Bills, which helps customers better understand usage charges incurred over the billing period.\n\nSee [Working with Bill Statements](https://www.m3ter.com/docs/guides/running-viewing-and-managing-bills/working-with-bill-statements) in the m3ter documentation for more details.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *AccountsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
