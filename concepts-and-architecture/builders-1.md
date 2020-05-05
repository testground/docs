---
description: What are Testground builders?
---

# Builders

A `builder` is a process which compiles the code of your test plan into a work unit, ready to be used by the ground service. The build process will be different depending on the language of the plan and the kind of work unit being targeted.

```text
                      ☟
-------------    -----------    ----------------    ----------    ---------------
| plan code | -> | builder | -> | unit of work | -> | runner | -> | test output |
-------------    -----------    ----------------    ----------    ---------------
```

### Supported builders

| builder | input language | output type | compatible runners |
| :--- | :--- | :--- | :--- |
| exec:go | go | os-specific executable | local:exec |
| docker:go | go | docker image | local:docker, cluster:k8s |

## Builder flags

The builders accept flags on the command-line which can modify their behaviour. Each builder has a different set of configurable options. The chart below shows options for each of the builders.

### exec:go builder

The `exec:go` builder uses the user's own go installation to compile and build a binary. Using this builder will use and alter the user's go pkg cache. None of these are required and need only be edited if the defaults do not work well in your environment.‌

​[github](https://github.com/ipfs/testground/blob/master/pkg/build/golang/exec.go#L28)​

| parameter | explanation |
| :--- | :--- |
| module\_path | use an alternative gomod path |
| exec\_pkg | specify the package name |
| fresh\_gomod | remove and recreate `go.mod` files |

### docker:go builder

The `docker:go` builder uses the user's local Docker daemon to construct a Docker image. By default, the `docker:go` builder will leverage a `goproxy` container to speed up fetching of Go modules. Additionally, all builds are performed on an isolated Docker network.‌

​[github](https://github.com/ipfs/testground/blob/master/pkg/build/golang/docker.go#L40)​

| parameter | explanation |
| :--- | :--- |
| go\_version | override the version of Go used to compile the plan |
| module\_path | use an alternative gomod path in the container |
| exec\_pkg | specify the package name |
| fresh\_gomod | remove and recreate `go.mod` files |
| push\_registry | after build, push docker image to a remote registry |
| registry\_type | must be set if push\_registry is true. Set to `aws` or `dockerhub` |
| go\_proxy\_mode | how to access go proxy. By default, use a local container. |
| go\_proxy\_url | required if `go_proxy_mode` is custom. Use a custom go\_proxy instance. |

## Examples

Single build for a single test for the example/output plan using the exec:go builder. This command will produce a binary which you can find in `~/testground/` on Linux and macOS systems.

```bash
$ testground build single --plan=example --builder=exec:go
```

Same, using the `docker:go` builder. This command will produce a docker image.

```bash
$ testground build single --plan=example --builder=docker:go
```

Use the `docker:go` builder to build an image and then upload the image to DockerHub.

```bash
$ testground build single --plan=example --builder=docker:go \
                                         --build-cfg push_registry=true \
                                         --build-cfg registry_type=dockerhub
```

Build a composition defined in `barrier-local.toml`. Note that the composition file will contain the builder and runner so specifying the builder on the command-line is not used in this example.

```bash
$ testground build composition --file=compositions/barrier-local.toml \
                               --write-artifacts
```

