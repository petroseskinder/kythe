load("//tools:build_rules/shims.bzl", "go_test")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "dedup",
    srcs = ["dedup.go"],
)

go_test(
    name = "dedup_test",
    size = "small",
    srcs = ["dedup_test.go"],
    library = "dedup",
    visibility = ["//visibility:private"],
    deps = ["//kythe/go/test/testutil"],
)
