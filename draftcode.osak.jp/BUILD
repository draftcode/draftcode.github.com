load("//rules:def.bzl", "remapped_filegroup", "render_single_page", "render_articles", "static_files")

remapped_filegroup(
    name = "final_layout",
    files = {
        ":static": "",
        ":computer_memo.html": "others/computer_memo.html",
        ":index.html": "index.html",
        ":contents": "contents",
    },
)

static_files(
    name = "static",
    srcs = glob(["static/**/*"]),
    trim_prefix = "static",
)

render_single_page(
    name = "index.html",
    src = "pages/index.md",
)

render_single_page(
    name = "computer_memo.html",
    src = "pages/computer_memo.md",
)

render_articles(
    name = "contents",
    srcs = glob(["posts/*"]),
)
