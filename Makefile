.PHONY: build-cmds
build-cmds:
	bazel build //cmd/...

.PHONY: build-functions
build-functions:
	bazel build --config linux //function/...

.PHONY: protoc-go
protoc-go:
	protoc --go_out=paths=source_relative:. types/v1/*.proto

.PHONY: update-bazel-deps
update-bazel-deps:
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: test
test:
	bazel test //...
