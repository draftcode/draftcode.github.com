load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "single_page_renderer_lib",
    srcs = ["single_page_renderer.go"],
    importpath = "github.com/draftcode/draftcode.github.com/rules/renderer/single_page_renderer",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_yuin_goldmark//:goldmark",
        "@com_github_yuin_goldmark//extension",
        "@com_github_yuin_goldmark//parser",
        "@com_github_yuin_goldmark//renderer/html",
        "@com_github_yuin_goldmark_highlighting//:goldmark-highlighting",
        "@com_github_yuin_goldmark_meta//:goldmark-meta",
    ],
)

go_binary(
    name = "single_page_renderer",
    embed = [":single_page_renderer_lib"],
    visibility = ["//visibility:public"],
)
