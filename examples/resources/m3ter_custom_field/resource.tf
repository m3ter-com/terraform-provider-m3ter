resource "m3ter_custom_field" "example_custom_field" {
  org_id = "orgId"
  account = {
    foo = "string"
  }
  account_plan = {
    foo = "string"
  }
  aggregation = {
    foo = "string"
  }
  compound_aggregation = {
    foo = "string"
  }
  contract = {
    foo = "string"
  }
  meter = {
    foo = "string"
  }
  organization = {
    foo = "string"
  }
  plan = {
    foo = "string"
  }
  plan_template = {
    foo = "string"
  }
  product = {
    foo = "string"
  }
}

#Singleton resources must be imported
import {
  to = m3ter_custom_field.example_custom_field
  id = "orgId"
}
