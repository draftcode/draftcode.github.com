def path_mapping(src, dst):
    return struct(src = src, dst = dst).to_json()

def path_mappings(mappings):
    return "[" + ",".join(mappings) + "]"

def _remapped_filegroup(ctx):
    mappings = []
    files = []
    for target, m in ctx.attr.files.items():
        mappings.append('{{"package": "{}", "mappings": {}}}'.format(target.label.package, m))
        files += target.files.to_list()

    mapping_file = ctx.actions.declare_file(ctx.label.name + "_mappings")
    ctx.actions.write(
        output = mapping_file,
        content = "[" + ",".join(mappings) + "]",
    )

    output = ctx.actions.declare_directory(ctx.label.name)
    args = ctx.actions.args()
    args.add("--mapping_file", mapping_file)
    args.add("--output_dir", output.path)
    ctx.actions.run(
        outputs = [output],
        executable = ctx.file._remapper,
        inputs = files + [mapping_file],
        arguments = [args],
    )
    return [DefaultInfo(files = depset([output]))]

remapped_filegroup = rule(
    attrs = {
        "files": attr.label_keyed_string_dict(
            mandatory = True,
        ),
        "_remapper": attr.label(
            default = "//rules/remapper",
            allow_single_file = True,
            executable = True,
            cfg = "exec",
        ),
    },
    implementation = _remapped_filegroup,
)

def _gcs_upload(ctx):
    command = """#!/bin/sh
    gsutil -m rsync -r "{}" "{}"
    """.format(ctx.file.src.short_path, ctx.attr.dst)
    ctx.actions.write(
        output = ctx.outputs.executable,
        content = command,
    )
    return [DefaultInfo(runfiles = ctx.runfiles(files = [ctx.file.src]))]

gcs_upload = rule(
    attrs = {
        "dst": attr.string(
            mandatory = True,
        ),
        "src": attr.label(
            mandatory = True,
            allow_single_file = True,
        ),
    },
    executable = True,
    implementation = _gcs_upload,
)
