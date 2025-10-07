# M3ter Terraform Provider

The [M3ter Terraform provider](https://registry.terraform.io/providers/m3ter-com/m3ter/latest/docs) provides convenient access to
the [M3ter REST API](https://www.m3ter.com) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    m3ter = {
      source  = "m3ter-com/m3ter"
      version = "~> 0.6.0"
    }
  }
}

# Initialize the provider
provider "m3ter" {
  api_key = "My API Key" # or set M3TER_API_KEY env variable
  api_secret = "My API Secret" # or set M3TER_API_SECRET env variable
  token = "My Token" # or set M3TER_API_TOKEN env variable
  org_id = "My Org ID" # or set M3TER_ORG_ID env variable
}

# Configure a resource
resource "m3ter_product" "example_product" {
  org_id = "My Org ID"
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/m3ter-com/m3ter/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property   | Environment variable | Required | Default value |
| ---------- | -------------------- | -------- | ------------- |
| org_id     | `M3TER_ORG_ID`       | true     | —             |
| api_secret | `M3TER_API_SECRET`   | true     | —             |
| api_key    | `M3TER_API_KEY`      | true     | —             |
| token      | `M3TER_API_TOKEN`    | false    | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/m3ter-com/terraform-provider-m3ter/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
