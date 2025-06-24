resource "m3ter_plan" "example_plan" {
  org_id = "orgId"
  code = "JS!?Q0]r] ]$]"
  name = "x"
  plan_template_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  account_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  bespoke = true
  custom_fields = {
    foo = "string"
  }
  minimum_spend = 0
  minimum_spend_accounting_product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  minimum_spend_bill_in_advance = true
  minimum_spend_description = "minimumSpendDescription"
  ordinal = 0
  standing_charge = 0
  standing_charge_accounting_product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  standing_charge_bill_in_advance = true
  standing_charge_description = "standingChargeDescription"
  version = 0
}
