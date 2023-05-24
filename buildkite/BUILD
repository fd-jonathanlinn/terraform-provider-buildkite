load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "buildkite",
    srcs = [
        "client.go",
        "data_source_meta.go",
        "data_source_organization.go",
        "data_source_pipeline.go",
        "data_source_team.go",
        "generated.go",
        "provider.go",
        "resource_agent_token.go",
        "resource_organization_settings.go",
        "resource_pipeline.go",
        "resource_pipeline_schedule.go",
        "resource_team.go",
        "resource_team_member.go",
        "util.go",
    ],
    importpath = "github.com/buildkite/terraform-provider-buildkite/buildkite",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_hashicorp_terraform_plugin_sdk_v2//diag",
        "@com_github_hashicorp_terraform_plugin_sdk_v2//helper/schema",
        "@com_github_khan_genqlient//graphql",
        "@com_github_shurcool_graphql//:graphql",
    ],
)

go_test(
    name = "buildkite_test",
    srcs = [
        "data_source_meta_test.go",
        "data_source_pipeline_test.go",
        "data_source_team_test.go",
        "provider_test.go",
        "resource_agent_token_test.go",
        "resource_organization_settings_test.go",
        "resource_pipeline_schedule_test.go",
        "resource_pipeline_test.go",
        "resource_team_member_test.go",
        "resource_team_test.go",
        "util_test.go",
    ],
    embed = [":buildkite"],
    deps = [
        "@com_github_hashicorp_terraform_plugin_sdk_v2//diag",
        "@com_github_hashicorp_terraform_plugin_sdk_v2//helper/resource",
        "@com_github_hashicorp_terraform_plugin_sdk_v2//helper/schema",
        "@com_github_hashicorp_terraform_plugin_sdk_v2//terraform",
    ],
)
