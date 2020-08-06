---
description: Building plans
---

# Builders

A **builder** is a component that takes a plan source code, and optionally the SDK's source code, and compiles it into a _**build artifact**_, ready to be used by compatible [Runners](runners.md) to schedule test workloads on Testground.

The build process is different depending on the language of the plan, and the kind of build artifact being targeted. Here's a simple diagram to understand how builders and runners relate to one another.

```text
                      ☟
-------------    -----------    ------------------    ----------    ---------------
| plan code | -> | builder | -> | build artifact | -> | runner | -> | test output |
-------------    -----------    ------------------    ----------    ---------------
```

## Supported builders

Builder names follow the format: `<build artifact type>:<language>`

| builder | input language | output type | compatible runners |
| :--- | :--- | :--- | :--- |
| `exec:go` | Go | OS-specific executable | `local:exec` |
| `docker:go` | Go | Docker image | `local:docker`, `cluster:k8s` |
| `docker:generic` | Any | Docker image | `local:docker`, `cluster:k8s` |

## Builder configuration options

The builders accept options on the command-line which can customize their behaviour. Each builder has a different set of configurable options. This section lists the configuration options supported by each.

Builder configuration options can be provided by various means, in the following order of precedence \(highest precedence to lowest precedence\):

1. CLI `--build-cfg` flags for `single` commands, and in the composition file for `composition` commands.
2.  `.env.toml`: `[builders]` section.
3. Test plan manifest.
4. Builder defaults \(applied by the runner\).

### exec:go builder

The `exec:go` builder uses the machine's own Go installation to compile and build a binary. Below are the options this builder supports. None of these are required and need only be edited if the defaults do not work well in your environment.‌

| options | explanation |
| :--- | :--- |
| `module_path` | gomod path with `fresh_gomod` |
| `exec_pkg` | specify the package to build |
| `fresh_gomod` | remove and recreate `go.mod` files |

### docker:go builder

The `docker:go` builder uses the user's local Docker daemon to construct a Docker image. By default, the `docker:go` builder will leverage a `goproxy` container to speed up fetching of Go modules. Additionally, all builds are performed on an isolated Docker network.‌

None of these options are required and need only be edited if the defaults do not work well in your environment.‌

| option | explanation |
| :--- | :--- |
| `go_version` | override the version of Go used to compile the plan |
| `module_path` | gomod path with `fresh_gomod` |
| `exec_pkg` | specify the package to build |
| `fresh_gomod` | remove and recreate `go.mod` files |
| `push_registry` | after build, push docker image to a remote registry |
| `registry_type` | must be set if push\_registry is true. Set to `aws` or `dockerhub` |
| `go_proxy_mode` | how to access go proxy. By default, use a local container. |
| `go_proxy_url` | required if `go_proxy_mode` is custom. Use a custom go\_proxy instance. |

## Examples

Single build for a single test for the example/output plan using the exec:go builder. This command will produce a binary which you can find in `~/testground/` on Linux and macOS systems.

```bash
$ testground build single --plan=example --builder=exec:go
```

Same, using the `docker:go` builder. This command will produce a Docker image.

```bash
$ testground build single --plan=example --builder=docker:go
```

Use the `docker:go` builder to build an image and then push the image to DockerHub \(configure credentials in [env.toml file](../getting-started.md)\).

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

