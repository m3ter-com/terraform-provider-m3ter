---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "m3ter_product Resource - m3ter"
subcategory: ""
description: |-
  
---

# m3ter_product (Resource)



## Example Usage

```terraform
resource "m3ter_product" "example_product" {
  org_id = "orgId"
  code = "S?oC\"$]C] ]]]]]5]"
  name = "x"
  custom_fields = {
    foo = "string"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `code` (String) A unique short code to identify the Product. It should not contain control chracters or spaces.
- `name` (String) Descriptive name for the Product providing context and information.

### Optional

- `custom_fields` (Dynamic) User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.

If `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.

See [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.
- `org_id` (String, Deprecated)

### Read-Only

- `created_by` (String) The unique identifier (UUID) of the user who created this Product.
- `dt_created` (String) The date and time *(in ISO-8601 format)* when the Product was created.
- `dt_last_modified` (String) The date and time *(in ISO-8601 format)* when the Product was last modified.
- `id` (String) The UUID of the entity.
- `last_modified_by` (String) The unique identifier (UUID) of the user who last modified this Product.
- `version` (Number) The version number:
- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.
- **Update:** On successful Update, the version is incremented by 1 in the response.

## Import

Import is supported using the following syntax:

```shell
$ terraform import m3ter_product.example '<org_id>/<id>'
```
