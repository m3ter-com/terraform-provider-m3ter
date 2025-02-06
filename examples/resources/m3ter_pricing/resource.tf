resource "m3ter_pricing" "example_pricing" {
  org_id = "orgId"
  pricing_bands = [{
    fixed_price = 0
    lower_limit = 0
    unit_price = 0
    id = "id"
    credit_type_id = "creditTypeId"
  }]
  start_date = "2019-12-27T18:11:19.117Z"
  aggregation_id = "aggregationId"
  code = "JS!?Q0]r] ]$]"
  compound_aggregation_id = "compoundAggregationId"
  cumulative = true
  description = "description"
  end_date = "2019-12-27T18:11:19.117Z"
  minimum_spend = 0
  minimum_spend_bill_in_advance = true
  minimum_spend_description = "minimumSpendDescription"
  overage_pricing_bands = [{
    fixed_price = 0
    lower_limit = 0
    unit_price = 0
    id = "id"
    credit_type_id = "creditTypeId"
  }]
  plan_id = "planId"
  plan_template_id = "planTemplateId"
  segment = {
    foo = "string"
  }
  tiers_span_plan = true
  type = "DEBIT"
  version = 0
}
