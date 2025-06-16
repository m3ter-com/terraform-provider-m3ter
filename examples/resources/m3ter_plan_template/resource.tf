resource "m3ter_plan_template" "example_plan_template" {
  org_id = "orgId"
  bill_frequency = "DAILY"
  currency = "xxx"
  name = "x"
  product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standing_charge = 0
  bill_frequency_interval = 1
  code = "JS!?Q0]r] ]$]"
  custom_fields = {
    foo = "string"
  }
  minimum_spend = 0
  minimum_spend_bill_in_advance = true
  minimum_spend_description = "minimumSpendDescription"
  ordinal = 0
  standing_charge_bill_in_advance = true
  standing_charge_description = "standingChargeDescription"
  standing_charge_interval = 1
  standing_charge_offset = 0
}
