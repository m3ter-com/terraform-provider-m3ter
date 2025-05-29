resource "m3ter_statement_statement_definition" "example_statement_statement_definition" {
  org_id = "orgId"
  aggregation_frequency = "DAY"
  dimensions = [{
    filter = ["string"]
    name = "x"
    attributes = ["string"]
    meter_id = "meterId"
  }]
  include_price_per_unit = true
  measures = [{
    aggregations = ["SUM"]
    meter_id = "meterId"
    name = "name"
  }]
  name = "name"
  version = 0
}
