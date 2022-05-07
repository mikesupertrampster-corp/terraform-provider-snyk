terraform {
  required_providers {
    snyk = {
      source = "localhost/mikesupertrampster-corp/snyk"
    }
  }
}

provider "snyk" {}

data "snyk_project" "one" {
  name = "mikesupertrampster-corp/blockchain:simple/go.mod"
}

output "test" {
  value = data.snyk_project.one
}