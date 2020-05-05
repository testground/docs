---
description: Running test plans
---

# What are Runners?

A `runner` is system which executes a number of instances of plans which have previously been compiled by the Testground builder. Whereas the `builder` takes plan code as its input and produces the compiled unit of work as output, the work unit is executed by the `runner` with the right number of duplicates to generate the test output.

```text
                                                        ☟
-------------    -----------    ----------------    ----------    ---------------
| plan code | -> | builder | -> | unit of work | -> | runner | -> | test output |
-------------    -----------    ----------------    ----------    ---------------
```

### Supported Runners

| builder | input work unit | Environment |
| :--- | :--- | :--- |
| local:exec | os-specific executable | local |
| local:docker | docker image | local docker server |
| cluster:k8s | docker image | kubernetes |

## Runner flags

The builder and runner accept flags on the command-line which can modify their behaviour. The chart below shows available options for each runner.

### local:exec runner

[github](https://github.com/ipfs/testground/blob/master/pkg/runner/local_exec.go#L42)

This runner currently has no configuration options.

### local:docker runner

[github](https://github.com/ipfs/testground/blob/master/pkg/runner/local_docker.go#L49)

| parameter | explanation |
| :--- | :--- |
| keep\_containers | specify whether containers should be removed after execution |
| log\_level | specify the logging verbosity |
| no\_start | if set, containers will be created but not executed |
| background | if set, the output of containers will not be displayed |
| ulimits | override ulimits applied to docker containers |

### cluster:k8s runner

[github](https://github.com/ipfs/testground/blob/master/pkg/runner/cluster_k8s.go#L120)

| parameter | explanation |
| :--- | :--- |
| kubeConfigPath | the location of your Kubernetes configuration file, if not in the default location |
| namespace | the Kubernetes namespace where plan pods will be scheduled |

## Examples

Build and run example/output with a single instance on the `local:docker` runner

```text
$ testground run single --builder=docker:go \
                        --runner=local:docker \
                        --plan=example \
                        --testcase=output \
                        --instances=1
```

with the `local:exec` runner:

```text
$ testground run single --builder=exec:go \
                        --runner=local:exec \
                        --plan=example \
                        --testcase=output \
                        --instances=1
```

and on a Kubernetes cluster \(and push the image to an AWS ECR container registry\)

```text
$ testground run single --builder=docker:go \
                        --build-cfg push_registry=true \
                        --build-cfg registry_type=aws \
                        --runner=cluster:k8s \
                        --plan=example \
                        --testcase=output \
                        --instances=1
```

