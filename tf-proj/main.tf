terraform {
  required_providers {
    buildkite = {
      source  = "buildkite/buildkite"
      version = "0.17.1"
    }
  }
}

provider "buildkite" {
}

resource "buildkite_pipeline" "test_migration" {
  name       = "Testing GQL Migration"
  repository = "https://github.com/buildkite/terraform-provider-buildkite.git"
}

output "badge_url" {
  value = buildkite_pipeline.test_migration.badge_url
}
