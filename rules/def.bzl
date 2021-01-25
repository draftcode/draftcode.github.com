def _trim_prefix(s, prefix):
    if s.startswith(prefix):
        return s[len(prefix):]
    return s

def _static_files(ctx):
    output = ctx.actions.declare_directory(ctx.label.name)
    commands = ["#!/bin/bash"]
    inputs = []
    for src in ctx.attr.srcs:
        for f in src.files.to_list():
            dst = _trim_prefix(f.path, ctx.label.package + "/")
            if ctx.attr.trim_prefix:
                dst = _trim_prefix(dst, ctx.attr.trim_prefix + "/")
            dst = output.path + "/" + dst
            commands.append("mkdir -p $(dirname \"{}\")".format(dst))
            commands.append("cp \"{}\" \"{}\"".format(f.path, dst))
            inputs.append(f)

    mapping_file = ctx.actions.declare_file(ctx.label.name + "_copy_files")
    ctx.actions.write(
        output = mapping_file,
        content = "\n".join(commands),
        is_executable = True,
    )
    ctx.actions.run(
        outputs = [output],
        executable = mapping_file,
        inputs = inputs + [mapping_file],
    )

    return [DefaultInfo(files = depset([output]))]

static_files = rule(
    attrs = {
        "srcs": attr.label_list(
            allow_files = True,
            mandatory = True,
        ),
        "trim_prefix": attr.string(),
    },
    implementation = _static_files,
)

def _remapped_filegroup(ctx):
    mappings = []
    files = []
    for target, dst in ctx.attr.files.items():
        for f in target.files.to_list():
            mappings.append('{{"src": "{}", "dst": "{}"}}'.format(f.path, dst))
            files.append(f)

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

def _render_single_page(ctx):
    output = ctx.actions.declare_file(ctx.label.name)
    args = ctx.actions.args()
    args.add("--input_file", ctx.file.src)
    args.add("--output_file", output.path)
    ctx.actions.run(
        outputs = [output],
        executable = ctx.file._single_page_renderer,
        inputs = [ctx.file.src],
        arguments = [args],
    )
    return [DefaultInfo(files = depset([output]))]

render_single_page = rule(
    attrs = {
        "src": attr.label(
            allow_single_file = True,
            mandatory = True,
        ),
        "_single_page_renderer": attr.label(
            default = "//rules/renderer/single_page_renderer",
            allow_single_file = True,
            executable = True,
            cfg = "exec",
        ),
    },
    implementation = _render_single_page,
)
