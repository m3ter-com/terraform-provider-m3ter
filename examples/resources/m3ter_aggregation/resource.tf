resource "m3ter_aggregation" "example_aggregation" {
  org_id = "orgId"
  aggregation = "SUM"
  meter_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  name = "x"
  quantity_per_unit = 0
  rounding = "UP"
  target_field = "x"
  unit = "x"
  code = "{1{}}_"
  custom_fields = {
    foo = "bar"
  }
  default_value = 0
  segmented_fields = ["string"]
  segments = [{
    foo = "string"
  }]
  version = 0
}
