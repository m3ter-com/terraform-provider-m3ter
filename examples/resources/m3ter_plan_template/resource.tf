resource "m3ter_plan_template" "example_plan_template" {
  org_id = "orgId"
  bill_frequency = "DAILY"
  currency = "USD"
  name = "string"
  product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standing_charge = 0
  bill_frequency_interval = 1
  code = "string"
  custom_fields = {

  }
  minimum_spend = 0
  minimum_spend_bill_in_advance = true
  minimum_spend_description = "string"
  ordinal = 0
  standing_charge_bill_in_advance = true
  standing_charge_description = "string"
  standing_charge_interval = 1
  standing_charge_offset = 364
}
