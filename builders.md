---
description: Building plans
---

# What are builders?

A builder is a process which compiles the code of your test plan into a work unit, ready to be used by the ground service. The build process will be different depending on the language of the plan and the kind of work unit being targeted.

```
                      ☟
-------------    -----------    ----------------    ----------    ---------------
| plan code | -> | builder | -> | unit of work | -> | Runner | -> | test output |
-------------    -----------    ----------------    ----------    ---------------
```


### Supported builders:

| builder   | input language | output type            | compatible runners       |
|-----------|----------------|------------------------|--------------------------|
| exec:go   | go             | os-specific executable | local:exec               |
| docker:go | go             | docker image           | local:docker cluster:k8s |


# Builder flags
The builder and runner accept flags on the commandline which can modify their behavior. Each builder has a different set of configurable options. The chart below shows options for each of the builders.

### exec:go builder

The exec:go builder uses the user's own go installation to compile and build a binary. Using this builder will use and alter the user's go pkg cache. None of these are required and need only be edited if the defaults do not work well in your environment.

[github](https://github.com/ipfs/testground/blob/master/pkg/build/golang/exec.go#L28)

| parameter   | explanation                        |
|-------------|------------------------------------|
| module_path |  use an alternative gomod path     |
| exec_pkg    | Specify the package name           |
| fresh_gomod | Remove and recreate `go.mod` files |


### docker:go builder

The docker:go builder uses the user's local docker daemon to construct a docker image. By default, the local:docker builder will leverage a goproxy  container to speed up fetching go modules. Additionally, all builds are performed on an isolated docker network.

[github](https://github.com/ipfs/testground/blob/master/pkg/build/golang/docker.go#L40)

| parameter     | explanation                                                             |
|---------------|-------------------------------------------------------------------------|
| go_version    |  override the version of golang used to compile the plan                |
| module_path   | use an  alternative gomod path in the container                         |
| exec_pkg      | Specify the package name                                                |
| fresh_gomod   | Remove and recreate `go.mod` files                                      |
| push_registry | After build, push docker image to a remote registry                     |
| registry_type | must be set if push_registry is true. Set to `aws` or `hub`             |
| go_proxy_mode | how to access go proxy. By default, use a local container.              |
| go_proxy_url  | required if `go_proxy_mode` is custom. Use an custom go_proxy instance. |



# examples

Single build for a single test for the example/output plan using the exec:go builder. This command will produce a binary which you can find in `~/.testground/` on Linux and OSX systems.

```bash
./testground build single example/output --builder exec:go
```

Same, using the docker:go builder. This command will produce a docker image.

```bash
./testground build single example/output --builder docker:go
```

Use the docker builder to build an image and then upload the image to docker hub

```bash
./testground build single example/output --builder docker:go --build-cfg push_registry=true --build-cfg registry_type=hub
```

Build a composition defined in `barrier-local.toml`. Note that the composition file will contain the builder and runner so specifying the builder on the commandline is not used in this example. 

```bash
./testground build composition -f compositions/barrier-local.toml --write-artifacts
```
