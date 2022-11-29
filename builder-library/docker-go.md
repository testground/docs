# docker:go

## Background

Provided that you are writing plans in Go and using the [sync service sdk](https://github.com/testground/sdk-go),  `docker:go` is the builder you want to use. This builder produces plan artifacts which are usable on the `local:docker`  or `cluster:k8s` runner.

## Features

* Produces docker containers which can run on docker-powered runners.
* Linkage with a custom sync service SDK via `--link-sdk` flag
* Simple command-line dependency mapping using the `--dep` flag
* `busybox` base for easy container troubleshooting.
* customizable build

## Troubleshooting

Most build failures are caused by a problem with the plan code rather than the build process. Here are a few tips to help figure out what's going on when there are build failures.

1. Try to build the plan yourself. These are just executable files after all! Frequently, build problems can be revealed by simply trying to run the plan as it is. If you can't build it like this, then the builder won't be able to either.

   ```text
   cd <plan_directory>
   export GOMOD111MODULE=on go build .
   ```

2. Make sure you have correctly initialized `go.mod`. The builder may throw errors when doing module replacements not used correctly.
3. Remove any custom module replace lines you have in `go.mod`. The builder will do this for you with appropriate flags.

## Customizing the build

when using the `docker:go` builder, plans are build using a standard template. This template is typically all that is needed, but for some plans, the default docker build may be too inflexible. For cases such as this, the Dockerfile can be extended to include custom directives.

?> If you want to just provide your own Dockerfile, use the `docker:generic` builder instead.

This feature is best explained by showing how it works. This is the default Dockerfile the `docker:go` builder will use to build plans. Notice that this is a go template. The template has a few points where customizations can be added.

```text
ARG GO_VERSION=1.14.2
ARG RUNTIME_IMAGE=busybox:1.31.1-glibc
FROM golang:${GO_VERSION}-buster AS builder
ARG TESTPLAN_EXEC_PKG="."
ARG GO_PROXY=direct
ARG BUILD_TAGS
ENV TESTPLAN_EXEC_PKG ${TESTPLAN_EXEC_PKG}
ENV PLAN_DIR /plan/
COPY /plan/go.mod ${PLAN_DIR}

{{if .WithSDK}}
COPY /sdk/go.mod /sdk/go.mod
{{end}}

{{.DockerfileExtensions.PreModDownload}}

RUN echo "Using go proxy: ${GO_PROXY}" \
    && cd ${PLAN_DIR} \
    && go env -w GOPROXY="${GO_PROXY}" \
    && go mod download

{{.DockerfileExtensions.PostModDownload}}

{{.DockerfileExtensions.PreSourceCopy}}

COPY . /

{{.DockerfileExtensions.PostSourceCopy}}

{{.DockerfileExtensions.PreBuild}}

RUN cd ${PLAN_DIR} \
    && go env -w GOPROXY="${GO_PROXY}" \
    && GOOS=linux GOARCH=amd64 go build -o testplan ${BUILD_TAGS} ${TESTPLAN_EXEC_PKG}

{{.DockerfileExtensions.PostBuild}}

# Store module dependencies
RUN cd ${PLAN_DIR} \
  && go list -m all > /testground_dep_list


FROM ${RUNTIME_IMAGE} AS binary

{{.DockerfileExtensions.PreRuntimeCopy}}

COPY --from=builder /testground_dep_list /
COPY --from=builder /plan/testplan /

{{.DockerfileExtensions.PostRuntimeCopy}}

EXPOSE 6060
ENTRYPOINT [ "/testplan"]

```

To add additional directives, add a section to your plan's `manifest.toml`. This example will add echo statements to each of the templated sections.

```text
[builders."docker:go".dockerfile_extensions]
pre_mod_download    = "RUN echo 'at pre_mod_download'"
post_mod_download   = "RUN echo 'at post_mod_download'"
pre_source_copy     = "RUN echo 'at pre_source_copy'"
post_source_copy    = "RUN echo 'at post_source_copy'"
pre_build           = "RUN echo 'at pre_build'"
post_build          = "RUN echo 'at post_build'"
pre_runtime_copy    = "RUN echo 'at pre_runtime_copy'"
post_runtime_copy   = "RUN echo 'at post_runtime_copy'"
```

## Learn More

See an example plan which uses a customized Dockerfile, see [here](https://github.com/testground/testground/tree/master/plans/dockercustomize)

Interested in how this works? All the Testground builders can be seen [here](https://github.com/testground/testground/tree/master/pkg/build)
