load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "analysis",
    srcs = ["analysis.go"],
    deps = [
        "//kythe/proto:analysis_go_proto",
        "//kythe/proto:analysis_service_go_proto",
        "//kythe/proto:storage_go_proto",
        "@go_protobuf//:proto",
    ],
)
