terraform {
  required_providers {
    snyk = {
      source = "localhost/mikesupertrampster-corp/snyk"
    }
  }
}

provider "snyk" {}
