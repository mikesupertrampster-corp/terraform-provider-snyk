# Snyk Provider

## Example Usage

```hcl
# Terraform 0.13+ uses the Terraform Registry:

terraform {
  required_providers {
    snyk = {
      source = "localhost/mikesupertrampster-corp/snyk"
    }
  }
}

provider "snyk" {
  #  api_key  = "API_KEY"  # can also provide in env as SNYK_API_KEY
  #  org_id   = "ORG_ID"   # can also provide in env as SNYK_ORG_ID
}

data "snyk_projects" "all" {}

output "all_project_names" {
  value = data.snyk_projects.all
}
```
