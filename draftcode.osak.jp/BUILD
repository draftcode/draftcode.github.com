load("//rules:def.bzl", "remapped_filegroup", "render_single_page", "static_files")

static_files(
    name = "static",
    srcs = glob(["static/**/*"]),
    trim_prefix = "static",
)

remapped_filegroup(
    name = "final_layout",
    files = {
        ":static": "",
        ":computer_memo.html": "others/computer_memo.html",
    },
)

render_single_page(
    name = "computer_memo.html",
    src = "pages/computer_memo.md",
)
