load("//tools:build_rules/shims.bzl", "go_test")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "process",
    srcs = ["async.go"],
)

go_test(
    name = "process_test",
    size = "small",
    srcs = ["process_test.go"],
    library = "process",
    visibility = ["//visibility:private"],
    deps = ["//kythe/go/test/testutil"],
)
