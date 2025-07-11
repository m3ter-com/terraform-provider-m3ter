---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_permission_policies Data Source - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_permission_policies (Data Source)



## Example Usage

```terraform
data "m3ter_permission_policies" "example_permission_policies" {
  org_id = "orgId"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `max_items` (Number) Max items to fetch, default: 1000
- `org_id` (String, Deprecated)

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `created_by` (String) The unique identifier (UUID) of the user who created this Permission Policy.
- `dt_created` (String) The date and time *(in ISO-8601 format)* when the Permission Policy was created.
- `dt_last_modified` (String) The date and time *(in ISO-8601 format)* when the Permission Policy was last modified.
- `id` (String) The unique identifier (UUID) for this Permission Policy.
- `last_modified_by` (String) The unique identifier (UUID) of the user who last modified this Permission Policy.
- `managed_policy` (Boolean) Indicates whether this is a system generated Managed Permission Policy.
- `name` (String) The name of the Permission Policy.
- `permission_policy` (Attributes List) Array containing the Permission Policies information. (see [below for nested schema](#nestedatt--items--permission_policy))
- `version` (Number) The version number. Default value when newly created is one.

<a id="nestedatt--items--permission_policy"></a>
### Nested Schema for `items.permission_policy`

Read-Only:

- `action` (List of String) The actions available to users who are assigned the Permission Policy - what they can do or cannot do with respect to the specified resource.

**NOTE:** Use lower case and a colon-separated format, for example, if you want to confer full CRUD, use:
```
"config:create",
"config:delete",
"config:retrieve",
"config:update"
```
- `effect` (String) Specifies whether or not the user is allowed to perform the action on the resource.

**NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll receive an error.
Available values: "ALLOW", "DENY".
- `resource` (List of String) See [Statements - Available Resources](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions#statements---available-resources) for a listing of available resources for Permission Policy statements.
