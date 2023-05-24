load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/buildkite/terraform-provider-buildkite
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "terraform-provider-buildkite_lib",
    srcs = ["main.go"],
    importpath = "github.com/buildkite/terraform-provider-buildkite",
    visibility = ["//visibility:private"],
    deps = [
        "//buildkite",
        "@com_github_hashicorp_terraform_plugin_sdk_v2//helper/schema",
        "@com_github_hashicorp_terraform_plugin_sdk_v2//plugin",
    ],
)

go_binary(
    name = "terraform-provider-buildkite",
    embed = [":terraform-provider-buildkite_lib"],
    visibility = ["//visibility:public"],
)
