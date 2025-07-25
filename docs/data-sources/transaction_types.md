---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_transaction_types Data Source - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_transaction_types (Data Source)



## Example Usage

```terraform
data "m3ter_transaction_types" "example_transaction_types" {
  org_id = "orgId"
  archived = true
  codes = ["string"]
  ids = ["string"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `archived` (Boolean) Filter with this Boolean flag whether to include TransactionTypes that are archived. 

* TRUE - include archived TransactionTypes in the list.
* FALSE - exclude archived TransactionTypes.
- `codes` (List of String) A list of TransactionType short codes to retrieve.
- `ids` (List of String) A list of TransactionType unique identifiers (UUIDs) to retrieve.
- `max_items` (Number) Max items to fetch, default: 1000
- `org_id` (String, Deprecated)

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `archived` (Boolean) TRUE / FALSE flag indicating whether the data entity is archived. An entity can be archived if it is obsolete.
- `code` (String) The short code of the data entity.
- `created_by` (String) The unique identifier (UUID) of the user who created this TransactionType.
- `dt_created` (String) The date and time *(in ISO-8601 format)* when the TransactionType was created.
- `dt_last_modified` (String) The date and time *(in ISO-8601 format)* when the TransactionType was last modified.
- `id` (String) The UUID of the entity.
- `last_modified_by` (String) The unique identifier (UUID) of the user who last modified this TransactionType.
- `name` (String) The name of the data entity.
- `version` (Number) The version number:
- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.
- **Update:** On successful Update, the version is incremented by 1 in the response.
