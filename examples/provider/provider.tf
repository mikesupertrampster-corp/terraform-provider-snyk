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