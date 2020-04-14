---
description: Running plans
---

# What are runners?

A runner is system which executes a number of instances of plans which have previously been compiled by the testground builder. Whereas the builder takes plan code as its input and produces the compiled unit of work as output, the work unit and executes it with the right number of duplicates to generate the test output.

                                                        ☟
-------------    -----------    ----------------    ----------    ---------------
| plan code | -> | builder | -> | unit of work | -> | Runner | -> | test output |
-------------    -----------    ----------------    ----------    ---------------


### Supported Runners

|-------------------------------------------------------------|
| builder      | input work unit        | Environment         | 
|_____________________________________________________________|
| local:exec   | os-specific executable | local               |
|--------------|------------------------|---------------------|
| local:docker | docker image           | local docker server |
|--------------|------------------------|---------------------|
| cluster:k8s  | docker image           | kubernetes          |
|-------------------------------------------------------------|


# Runner flags
The builder and runner accept flags on the commandline which can modify their behavior. The chart below shows available options for each runner.

### local:exec runner

[github](https://github.com/ipfs/testground/blob/master/pkg/runner/local_exec.go#L42)

This runner currently has no configuration options.

### local:docker runner
[github](https://github.com/ipfs/testground/blob/master/pkg/runner/local_docker.go#L49)

|---------------------------------------------------------------------------------|
| parameter       | explanation                                                   |
|_________________________________________________________________________________|
| keep_containers |  Specify whether containers should be removed after execution |
|---------------------------------------------------------------------------------|
| log_level       | Specify the logginv verbosity                                 |
|---------------------------------------------------------------------------------|
| no_start        | if set, containers will be created but not executed.          |
|---------------------------------------------------------------------------------|
| background      | if set, the output of containers will not be displayed.       |
|---------------------------------------------------------------------------------|
| Ulimits         | Override ulimits applied to docker containers.                |
|---------------------------------------------------------------------------------|

### cluster:k8s runner
[github](https://github.com/ipfs/testground/blob/master/pkg/runner/cluster_k8s.go#L120)

|------------------------------------------------------------------------------------------------------|
| parameter      | explanation                                                                         |
|______________________________________________________________________________________________________|
| kubeConfigPath |  The location of your kubernetes configuration file, if not in the defualt location |
|------------------------------------------------------------------------------------------------------|
| namespace      | The kubernetes namespace where plan pods will be scheduled                          |
|------------------------------------------------------------------------------------------------------|


# examples

build and run example/output with a single instance on the local:docker runner
```
./testground run single --builder docker:go --runner local:docker example/output -i 1

```

with the local:exec runner:

```
./testground run single --builder exec:go --runner local:exec example/output -i 1

```

and on a kubernetes cluster (and push the image to an aws container registry)

```
./testground run single --builder docker:go --build-cfg push_registry=true --build-cfg registry_type=aws --runner cluster:k8s example/output -i 1
```
