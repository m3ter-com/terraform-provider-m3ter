---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_bill_config Resource - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_bill_config (Resource)



## Example Usage

```terraform
resource "m3ter_bill_config" "example_bill_config" {
  org_id = "orgId"
  bill_lock_date = "2019-12-27"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bill_lock_date` (String) The global lock date when all Bills will be locked *(in ISO 8601 format)*.

For example: `"2024-03-01"`.
- `org_id` (String, Deprecated)

### Read-Only

- `created_by` (String) The id of the user who created this bill config.
- `dt_created` (String) The DateTime *(in ISO-8601 format)* when the bill config was first created.
- `dt_last_modified` (String) The DateTime *(in ISO-8601 format)* when the bill config was last modified.
- `id` (String) The Organization UUID. The Organization represents your company as a direct customer of the m3ter service.
- `last_modified_by` (String) The id of the user who last modified this bill config.
- `version` (Number) The version number:
* Default value when newly created is one.
* Incremented by 1 each time it is updated.

## Import

Import is supported using the following syntax:

```shell
$ terraform import m3ter_bill_config.example '<org_id>'
```
