load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "single_page_renderer_lib",
    srcs = ["single_page_renderer.go"],
    importpath = "github.com/draftcode/draftcode.github.com/rules/renderer/single_page_renderer",
    visibility = ["//visibility:private"],
    deps = ["//rules/renderer"],
)

go_binary(
    name = "single_page_renderer",
    embed = [":single_page_renderer_lib"],
    visibility = ["//visibility:public"],
)
