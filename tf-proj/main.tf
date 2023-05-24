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

resource "buildkite_pipeline" "test_genqlient" {
  name       = "Testing Genqlient"
  repository = "https://github.com/buildkite/terraform-provider-buildkite.git"

  tags = [ "terraform", "buildkite", "genqlient" ]
}

output "badge_url" {
  value = buildkite_pipeline.test_genqlient.badge_url
}
