load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_go_playground_assert",
        importpath = "github.com/go-playground/assert",
        sum = "h1:ad06XqC+TOv0nJWnbULSlh3ehp5uLuQEojZY5Tq8RgI=",
        version = "v1.2.1",
    )
