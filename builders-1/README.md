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

The builder and runner accept flags on the command-line which can modify their behaviour. Each builder has a different set of configurable options. The chart below shows options for each of the builders.

### exec:go builder



### docker:go builder



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

