data "snyk_project" "blockchain_simple" {
  name = "mikesupertrampster-corp/blockchain:simple/go.mod"
}

output "blockchain_simple" {
  value = data.snyk_project.blockchain_simple
}

data "snyk_projects" "all" {}

output "all_project_names" {
  value = data.snyk_projects.all
}