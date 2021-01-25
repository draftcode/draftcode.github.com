load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "remapper",
    embed = [":remapper_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "remapper_lib",
    srcs = ["remapper.go"],
    importpath = "github.com/draftcode/draftcode.github.com/rules/remapper",
    visibility = ["//visibility:private"],
)
