# terraform-provider-snyk

Terraform provider for configuring Snyk

[![Snyk Golang Scan](https://github.com/mikesupertrampster-corp/terraform-provider-snyk/actions/workflows/snyk.yml/badge.svg)](https://github.com/mikesupertrampster-corp/terraform-provider-snyk/actions/workflows/snyk.yml) [![gitleaks](https://github.com/mikesupertrampster-corp/terraform-provider-snyk/actions/workflows/gitleaks.yml/badge.svg)](https://github.com/mikesupertrampster-corp/terraform-provider-snyk/actions/workflows/gitleaks.yml) [![release](https://github.com/mikesupertrampster-corp/terraform-provider-snyk/actions/workflows/release.yml/badge.svg)](https://github.com/mikesupertrampster-corp/terraform-provider-snyk/actions/workflows/release.yml) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a36286278570406692de2ac036bc7a94)](https://www.codacy.com/gh/mikesupertrampster-corp/terraform-provider-snyk/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mikesupertrampster-corp/terraform-provider-snyk&amp;utm_campaign=Badge_Grade) 

## Requirements

   - [Terraform](https://www.terraform.io/downloads.html) 0.11.x

## Building The Provider

Clone repository:

```sh
$ git clone git@github.com:mikesupertrampster-corp/terraform-provider-snyk
```

Build the provider:

```sh
$ make build
```

Install the provider:

```sh
$ make install
```

## Using the provider

If you're building the provider, follow the instructions to install the provider, then run `terraform init` to initialize it.

Use the locally installed provider by copy and pasting this into your Terraform configuration

```hcl
terraform {
  required_providers {
    snyk = {
      source = "localhost/mikesupertrampster-corp/snyk"
    }
  }
}
```
