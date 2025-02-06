resource "m3ter_commitment" "example_commitment" {
  org_id = "orgId"
  account_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  amount = 1
  currency = "x"
  end_date = "2019-12-27"
  start_date = "2019-12-27"
  accounting_product_id = "accountingProductId"
  amount_first_bill = 0
  amount_pre_paid = 0
  bill_epoch = "2019-12-27"
  billing_interval = 1
  billing_offset = 0
  billing_plan_id = "billingPlanId"
  child_billing_mode = "PARENT_SUMMARY"
  commitment_fee_bill_in_advance = true
  commitment_fee_description = "commitmentFeeDescription"
  commitment_usage_description = "commitmentUsageDescription"
  contract_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  fee_dates = [{
    amount = 1
    date = "2019-12-27"
    service_period_end_date = "2019-12-27T18:11:19.117Z"
    service_period_start_date = "2019-12-27T18:11:19.117Z"
  }]
  line_item_types = ["STANDING_CHARGE"]
  overage_description = "overageDescription"
  overage_surcharge_percent = 0
  product_ids = ["string"]
  separate_overage_usage = true
  version = 0
}
