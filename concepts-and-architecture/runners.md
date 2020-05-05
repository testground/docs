---
description: What are Testground runners?
---

# Runners

A `runner` is system which executes a number of instances of plans which have previously been compiled by the Testground builder. Whereas the `builder` takes plan code as its input and produces the compiled unit of work as output, the work unit is executed by the `runner` with the right number of duplicates to generate the test output.

```text
                                                        ☟
-------------    -----------    ----------------    ----------    ---------------
| plan code | -> | builder | -> | unit of work | -> | runner | -> | test output |
-------------    -----------    ----------------    ----------    ---------------
```

### Supported runners

| builder | input work unit | Environment |
| :--- | :--- | :--- |
| local:exec | os-specific executable | local |
| local:docker | docker image | local docker server |
| cluster:k8s | docker image | Kubernetes |



