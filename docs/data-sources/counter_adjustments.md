---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_counter_adjustments Data Source - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_counter_adjustments (Data Source)



## Example Usage

```terraform
data "m3ter_counter_adjustments" "example_counter_adjustments" {
  org_id = "orgId"
  account_id = "accountId"
  counter_id = "counterId"
  date = "date"
  date_end = "dateEnd"
  date_start = "dateStart"
  end_date_end = "endDateEnd"
  end_date_start = "endDateStart"
  sort_order = "sortOrder"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) List CounterAdjustment items for the Account UUID.
- `counter_id` (String) List CounterAdjustment items for the Counter UUID.
- `date` (String) List CounterAdjustment items for the given date.
- `date_end` (String)
- `date_start` (String)
- `end_date_end` (String) Only include CounterAdjustments with end dates earlier than this date.
- `end_date_start` (String) Only include CounterAdjustments with end dates equal to or later than this date.
- `max_items` (Number) Max items to fetch, default: 1000
- `org_id` (String, Deprecated)
- `sort_order` (String) Sort order for the results

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `account_id` (String) The Account ID the CounterAdjustment was created for.
- `counter_id` (String) The ID of the Counter that was used to make the CounterAdjustment on the Account.
- `created_by` (String) The ID of the user who created this item.
- `date` (String) The date the CounterAdjustment was created for the Account *(in ISO-8601 date format)*.
- `dt_created` (String) The DateTime when this item was created *(in ISO-8601 format)*.
- `dt_last_modified` (String) The DateTime when this item was last modified *(in ISO-8601 format)*.
- `id` (String) The UUID of the entity.
- `last_modified_by` (String) The ID of the user who last modified this item.
- `purchase_order_number` (String) Purchase Order Number for the Counter Adjustment. *(Optional)*
- `value` (Number) Integer Value of the Counter that was used to make the CounterAdjustment.
- `version` (Number) The version number:
- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.
- **Update:** On successful Update, the version is incremented by 1 in the response.
