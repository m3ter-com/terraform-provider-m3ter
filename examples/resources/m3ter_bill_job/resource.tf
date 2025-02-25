resource "m3ter_bill_job" "example_bill_job" {
  org_id = "orgId"
  account_ids = ["string"]
  bill_date = "2019-12-27"
  bill_frequency_interval = 0
  billing_frequency = "DAILY"
  currency_conversions = [{
    from = "EUR"
    to = "USD"
    multiplier = 1.12
  }]
  day_epoch = "2019-12-27"
  due_date = "2019-12-27"
  external_invoice_date = "2019-12-27"
  last_date_in_billing_period = "2019-12-27"
  month_epoch = "2019-12-27"
  target_currency = "xxx"
  timezone = "UTC"
  version = 0
  week_epoch = "2019-12-27"
  year_epoch = "2019-12-27"
}
