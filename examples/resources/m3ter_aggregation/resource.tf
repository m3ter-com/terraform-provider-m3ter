resource "m3ter_aggregation" "example_aggregation" {
  org_id = "orgId"
  aggregation = "SUM"
  meter_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  name = "x"
  quantity_per_unit = 1
  rounding = "UP"
  target_field = "x"
  unit = "x"
  accounting_product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  code = "example_code"
  custom_fields = {
    foo = "string"
  }
  custom_sql = "customSql"
  default_value = 0
  segmented_fields = ["string"]
  segments = [{
    foo = "string"
  }]
}
