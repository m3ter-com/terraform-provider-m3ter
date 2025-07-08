resource "m3ter_meter" "example_meter" {
  org_id = "orgId"
  code = "S?oC\"$]C] ]]]]]5]"
  data_fields = [{
    category = "WHO"
    code = "{1{}}_"
    name = "x"
    unit = "x"
  }]
  derived_fields = [{
    category = "WHO"
    code = "{1{}}_"
    name = "x"
    unit = "x"
    calculation = "x"
  }]
  name = "x"
  custom_fields = {
    foo = "string"
  }
  group_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  version = 0
}
