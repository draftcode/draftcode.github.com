load("//rules:def.bzl", "path_mapping", "path_mappings", "remapped_filegroup")

filegroup(
    name = "static",
    srcs = glob(["static/**/*"]),
)

remapped_filegroup(
    name = "final_layout",
    files = {
        ":static": path_mappings([
            path_mapping("static", ""),
        ]),
    },
)
