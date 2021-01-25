load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "article_renderer_lib",
    srcs = ["article_renderer.go"],
    importpath = "github.com/draftcode/draftcode.github.com/rules/renderer/article_renderer",
    visibility = ["//visibility:private"],
    deps = ["//rules/renderer"],
)

go_binary(
    name = "article_renderer",
    embed = [":article_renderer_lib"],
    visibility = ["//visibility:public"],
)
