data "snyk_project" "blockchain_simple" {
  name = "mikesupertrampster-corp/blockchain:simple/go.mod"
}

output "blockchain_simple" {
  value = data.snyk_project.blockchain_simple
}

