resource "m3ter_counter_pricing" "example_counter_pricing" {
  org_id = "orgId"
  counter_id = "x"
  pricing_bands = [{
    fixed_price = 0
    lower_limit = 0
    unit_price = 0
    id = "id"
    credit_type_id = "creditTypeId"
  }]
  start_date = "2019-12-27T18:11:19.117Z"
  accounting_product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  code = "S?oC\"$]C] ]]]]]5]"
  cumulative = true
  description = "description"
  end_date = "2019-12-27T18:11:19.117Z"
  plan_id = "planId"
  plan_template_id = "planTemplateId"
  pro_rate_adjustment_credit = true
  pro_rate_adjustment_debit = true
  pro_rate_running_total = true
  running_total_bill_in_advance = true
  version = 0
}
