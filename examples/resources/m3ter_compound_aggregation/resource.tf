resource "m3ter_compound_aggregation" "example_compound_aggregation" {
  org_id = "orgId"
  calculation = "x"
  name = "x"
  quantity_per_unit = 1
  rounding = "UP"
  unit = "x"
  code = "example_code"
  custom_fields = {
    foo = "string"
  }
  evaluate_null_aggregations = true
  product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  version = 0
}
