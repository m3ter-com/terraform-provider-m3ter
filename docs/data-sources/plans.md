---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_plans Data Source - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_plans (Data Source)



## Example Usage

```terraform
data "m3ter_plans" "example_plans" {
  org_id = "orgId"
  account_id = ["string"]
  ids = ["string"]
  product_id = "productId"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (List of String) List of Account IDs the Plan belongs to.
- `ids` (List of String) List of Plan IDs to retrieve.
- `max_items` (Number) Max items to fetch, default: 1000
- `org_id` (String, Deprecated)
- `product_id` (String) UUID of the Product to retrieve Plans for.

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `account_id` (String) *(Optional)*. The Account ID for which this plan was created as custom/bespoke. A custom/bespoke Plan can only be attached to the specified Account.
- `bespoke` (Boolean) TRUE/FALSE flag indicating whether the plan is custom/bespoke for a particular Account.
- `code` (String) Unique short code reference for the Plan.
- `created_by` (String) The id of the user who created this plan.
- `custom_fields` (Dynamic) User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.

If `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.

See [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.
- `dt_created` (String) The DateTime *(in ISO-8601 format)* when the plan was created.
- `dt_last_modified` (String) The DateTime *(in ISO-8601 format)* when the plan was last modified.
- `id` (String) The UUID of the entity.
- `last_modified_by` (String) The id of the user who last modified this plan.
- `minimum_spend` (Number) The product minimum spend amount per billing cycle for end customer Accounts on a priced Plan.

*(Optional)*. Overrides PlanTemplate value.
- `minimum_spend_accounting_product_id` (String) Optional Product ID this plan's minimum spend should be attributed to for accounting purposes
- `minimum_spend_bill_in_advance` (Boolean) When TRUE, minimum spend is billed at the start of each billing period.

When FALSE, minimum spend is billed at the end of each billing period.

*(Optional)*. Overrides the setting at PlanTemplate level for minimum spend billing in arrears/in advance.
- `minimum_spend_description` (String) Minimum spend description *(displayed on the bill line item)*.
- `name` (String) Descriptive name for the Plan.
- `ordinal` (Number) Assigns a rank or position to the Plan in your order of pricing plans - lower numbers represent more basic pricing plans; higher numbers represent more premium pricing plans.

*(Optional)*. Overrides PlanTemplate value.

**NOTE:** **DEPRECATED** - no longer used.
- `plan_template_id` (String) UUID of the PlanTemplate the Plan belongs to.
- `product_id` (String) UUID of the Product the Plan belongs to.
- `standing_charge` (Number) The standing charge applied to bills for end customers. This is prorated.

*(Optional)*. Overrides PlanTemplate value.
- `standing_charge_accounting_product_id` (String) Optional Product ID this plan's standing charge should be attributed to for accounting purposes
- `standing_charge_bill_in_advance` (Boolean) When TRUE, standing charge is billed at the start of each billing period.

When FALSE, standing charge is billed at the end of each billing period.

*(Optional)*. Overrides the setting at PlanTemplate level for standing charge billing in arrears/in advance.
- `standing_charge_description` (String) Standing charge description *(displayed on the bill line item)*.
- `version` (Number) The version number:
- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.
- **Update:** On successful Update, the version is incremented by 1 in the response.
