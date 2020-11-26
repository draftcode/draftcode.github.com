load("//rules:def.bzl", "gcs_upload", "path_mapping", "path_mappings", "remapped_filegroup")

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

gcs_upload(
    name = "upload",
    src = ":final_layout",
    dst = "gs://draftcode-osak-jp",
)
