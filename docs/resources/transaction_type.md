---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_transaction_type Resource - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_transaction_type (Resource)



## Example Usage

```terraform
resource "m3ter_transaction_type" "example_transaction_type" {
  org_id = "orgId"
  name = "x"
  archived = true
  code = "code"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the entity.

### Optional

- `archived` (Boolean) A Boolean TRUE / FALSE flag indicating whether the entity is archived. An entity can be archived if it is obsolete.

* TRUE - the entity is in the archived state.
* FALSE - the entity is not in the archived state.
- `code` (String) The short code for the entity.
- `org_id` (String, Deprecated)

### Read-Only

- `created_by` (String) The unique identifier (UUID) of the user who created this TransactionType.
- `dt_created` (String) The date and time *(in ISO-8601 format)* when the TransactionType was created.
- `dt_last_modified` (String) The date and time *(in ISO-8601 format)* when the TransactionType was last modified.
- `id` (String) The UUID of the entity.
- `last_modified_by` (String) The unique identifier (UUID) of the user who last modified this TransactionType.
- `version` (Number) The version number:
- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.
- **Update:** On successful Update, the version is incremented by 1 in the response.

## Import

Import is supported using the following syntax:

```shell
$ terraform import m3ter_transaction_type.example '<org_id>/<id>'
```
