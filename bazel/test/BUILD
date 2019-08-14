load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_to_json")

jsonnet_to_json(
    name = "output",
    src = "output.libsonnet",
    outs = ["output.json"],
    deps = [
        "//lib:lib"
    ],
)
