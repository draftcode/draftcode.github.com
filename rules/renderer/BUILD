load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "renderer",
    srcs = ["renderer.go"],
    importpath = "github.com/draftcode/draftcode.github.com/rules/renderer",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_yuin_goldmark//:goldmark",
        "@com_github_yuin_goldmark//extension",
        "@com_github_yuin_goldmark//parser",
        "@com_github_yuin_goldmark//renderer/html",
        "@com_github_yuin_goldmark_highlighting//:goldmark-highlighting",
        "@com_github_yuin_goldmark_meta//:goldmark-meta",
    ],
)
