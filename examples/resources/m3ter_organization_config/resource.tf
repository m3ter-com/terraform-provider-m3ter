resource "m3ter_organization_config" "example_organization_config" {
  org_id = "orgId"
  currency = "USD"
  day_epoch = "2022-01-01"
  days_before_bill_due = 1
  month_epoch = "2022-01-01"
  timezone = "UTC"
  week_epoch = "2022-01-04"
  year_epoch = "2022-01-01"
  auto_approve_bills_grace_period = 1
  auto_approve_bills_grace_period_unit = "DAYS"
  auto_generate_statement_mode = "NONE"
  bill_prefix = "Bill-"
  commitment_fee_bill_in_advance = true
  consolidate_bills = true
  credit_application_order = ["PREPAYMENT"]
  currency_conversions = [{
    from = "EUR"
    to = "USD"
    multiplier = 1
  }]
  default_statement_definition_id = "defaultStatementDefinitionId"
  external_invoice_date = "LAST_DAY_OF_ARREARS"
  minimum_spend_bill_in_advance = true
  scheduled_bill_interval = 0
  sequence_start_number = 1000
  standing_charge_bill_in_advance = true
  suppressed_empty_bills = true
  version = 0
}
