resource "m3ter_meter" "example_meter" {
  org_id = "orgId"
  code = "JS!?Q0]r] ]$]"
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
  }]
  name = "x"
  custom_fields = {
    foo = "bar"
  }
  group_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  version = 0
}
