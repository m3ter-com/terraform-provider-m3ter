---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_commitments Data Source - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_commitments (Data Source)



## Example Usage

```terraform
data "m3ter_commitments" "example_commitments" {
  org_id = "orgId"
  account_id = "accountId"
  contract_id = "contractId"
  date = "date"
  end_date_end = "endDateEnd"
  end_date_start = "endDateStart"
  ids = ["string"]
  product_id = "productId"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The unique identifier (UUID) for the Account. This parameter helps filter the Commitments related to a specific end-customer Account.
- `contract_id` (String)
- `date` (String) A date *(in ISO-8601 format)* to filter Commitments which are active on this specific date.
- `end_date_end` (String) A date *(in ISO-8601 format)* used to filter Commitments. Only Commitments with end dates before this date will be included.
- `end_date_start` (String) A date *(in ISO-8601 format)* used to filter Commitments. Only Commitments with end dates on or after this date will be included.
- `ids` (List of String) A list of unique identifiers (UUIDs) for the Commitments to retrieve. Use this to fetch specific Commitments in a single request.
- `max_items` (Number) Max items to fetch, default: 1000
- `org_id` (String, Deprecated)
- `product_id` (String) The unique identifier (UUID) for the Product. This parameter helps filter the Commitments related to a specific Product.

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `account_id` (String) The unique identifier (UUID) for the end customer Account the Commitment is added to.
- `accounting_product_id` (String) The unique identifier (UUID) for the Product linked to the Commitment for accounting purposes.
- `amount` (Number) The total amount that the customer has committed to pay.
- `amount_first_bill` (Number) The amount to be billed in the first invoice.
- `amount_pre_paid` (Number) The amount that the customer has already paid upfront at the start of the Commitment service period.
- `amount_spent` (Number) The total amount of the Commitment that the customer has spent so far.
- `bill_epoch` (String) The starting date *(in ISO-8601 date format)* from which the billing cycles are calculated.
- `billing_interval` (Number) How often the Commitment fees are applied to bills. For example, if the plan being used to bill for Commitment fees is set to issue bills every three months and the `billingInterval` is set to 2, then the Commitment fees are applied every six months.
- `billing_offset` (Number) The offset for when the Commitment fees are first applied to bills on the Account. For example, if bills are issued every three months and the `billingOffset` is 0, then the charge is applied to the first bill (at three months); if set to 1, it's applied to the next bill (at six months), and so on.
- `billing_plan_id` (String) The unique identifier (UUID) for the Product Plan used for billing Commitment fees due.
- `child_billing_mode` (String) If the Account is either a Parent or a Child Account, this specifies the Account hierarchy billing mode. The mode determines how billing will be handled and shown on bills for charges due on the Parent Account, and charges due on Child Accounts:

* **Parent Breakdown** - a separate bill line item per Account. Default setting.

* **Parent Summary** - single bill line item for all Accounts.

* **Child** - the Child Account is billed.
Available values: "PARENT_SUMMARY", "PARENT_BREAKDOWN", "CHILD".
- `commitment_fee_bill_in_advance` (Boolean) A boolean value indicating whether the Commitment fee is billed in advance *(start of each billing period)* or arrears *(end of each billing period)*.

* **TRUE** - bill in advance *(start of each billing period)*.
* **FALSE** - bill in arrears *(end of each billing period)*.
- `commitment_fee_description` (String) A textual description of the Commitment fee.
- `commitment_usage_description` (String) A textual description of the Commitment usage.
- `contract_id` (String) The unique identifier (UUID) for a Contract you've created for the Account and to which the Commitment has been added.
- `created_by` (String) The unique identifier (UUID) of the user who created this Commitment.
- `currency` (String) The currency used for the Commitment. For example, 'USD'.
- `drawdowns_accounting_product_id` (String)
- `dt_created` (String) The date and time *(in ISO-8601 format)* when the Commitment was created.
- `dt_last_modified` (String) The date and time *(in ISO-8601 format)* when the Commitment was last modified.
- `end_date` (String) The end date of the Commitment period in ISO-8601 format.
- `fee_dates` (Attributes List) Used for billing any outstanding Commitment fees *on a schedule*.

An array defining a series of bill dates and amounts covering specified service periods:
- `date` - the billing date *(in ISO-8601 format)*.
- `amount` - the billed amount.
- `servicePeriodStartDate` and `servicePeriodEndDate` - defines the service period the bill covers *(in ISO-8601 format)*. (see [below for nested schema](#nestedatt--items--fee_dates))
- `fees_accounting_product_id` (String)
- `id` (String) The UUID of the entity.
- `last_modified_by` (String) The unique identifier (UUID) of the user who last modified this Commitment.
- `line_item_types` (List of String) Specifies the line item charge types that can draw-down at billing against the Commitment amount. Options are:
- `MINIMUM_SPEND`
- `STANDING_CHARGE`
- `USAGE`
- `"COUNTER_RUNNING_TOTAL_CHARGE"`
- `"COUNTER_ADJUSTMENT_DEBIT"`
- `overage_description` (String) A textual description of the overage charges.
- `overage_surcharge_percent` (Number) The percentage surcharge applied to the usage charges that exceed the Commitment amount.
- `product_ids` (List of String) A list of unique identifiers (UUIDs) for Products the Account consumes. Charges due for these Products will be made available for draw-down against the Commitment.

**Note:** If not used, then charges due for all Products the Account consumes will be made available for draw-down against the Commitment.
- `separate_overage_usage` (Boolean) A boolean value indicating whether the overage usage is billed separately or together. If overage usage is separated and a Commitment amount has been consumed by an Account, any subsequent line items on Bills against the Account for usage will show as separate "overage usage" charges, not simply as "usage" charges:

* **TRUE** - billed separately.
* **FALSE** - billed together.
- `start_date` (String) The start date of the Commitment period in ISO-8601 format.
- `version` (Number) The version number:
- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.
- **Update:** On successful Update, the version is incremented by 1 in the response.

<a id="nestedatt--items--fee_dates"></a>
### Nested Schema for `items.fee_dates`

Read-Only:

- `amount` (Number)
- `date` (String)
- `service_period_end_date` (String)
- `service_period_start_date` (String)
