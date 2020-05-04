---
description: Writing your first test plan
---

# Quickstart

## Hello, Test Plans!

In this quick start tutorial you will get up and running with a simple test plan. Later tutorials will go deeper into features of the plan SDK and how to use it to simulate P2P network environments. But for now, let's get your hands dirty writing your first test plan. Hello!

### Create a plan manifest

Create a manifest in the `manifests` directory. This file is used to inform Testground about your plan.  This file describes the location of the plan, options specific to particular runners/builders, and any parameters that should be passed to the plan to control execution.

{% tabs %}
{% tab title="local docker" %}
{% code title="manifests/quickstart.toml" %}
```yaml
name = "quickstart"
source_path = "file://${TESTGROUND_SRCDIR}/plans/quickstart"

[defaults]
builder = "exec:go"
runner = "local:docker"

[build_strategies."docker:go"]
enabled = true
go_version = "1.13"
module_path = "github.com/ipfs/testground/plans/example"
exec_pkg = "."
go_ipfs_version = "0.4.22"

[run_strategies."local:docker"]
enabled = true

[[testcases]]
name = "testcase1"
instances = { min = 1, max = 200, default = 1 }
  [testcases.params]
  who = { type = "string", default="world" }
```
{% endcode %}
{% endtab %}

{% tab title="local exec" %}
{% code title="manifests/quickstart.toml" %}
```
name = "quickstart"
source_path = "file://${TESTGROUND_SRCDIR}/plans/quickstart"

[defaults]
builder = "exec:go"
runner = "local:exec"

[build_strategies."exec:go"]
enabled = true

[run_strategies."local:exec"]
enabled = true

[[testcases]]
name = "testcase1"
instances = { min = 1, max = 200, default = 1 }
  [testcases.params]
  who = { type = "string", default="world" }
```
{% endcode %}
{% endtab %}
{% endtabs %}

{% hint style="info" %}
You can enable multiple runners and builders in the same file! 
{% endhint %}

### Create a new test plan module 

```bash
$ mkdir -p plans/quickstart
$ pushd plans/quickstart
$ go mod init github.com/ipfs/testground/plans/quickstart
$ go mod edit -require "github.com/ipfs/testground/sdk/runtime@v0.3.0"
$ go mod edit -replace github.com/ipfs/testground/sdk/runtime=../../sdk/runtime
$ popd
```

### Write the plan

Fire up your favourite editor and input the following

{% code title="plans/quickstart/main.go" %}
```go
package main

import "github.com/ipfs/testground/sdk/runtime"

func main() {
    runtime.Invoke(run)
}

func run(runenv *runtime.RunEnv) error {
    runenv.RecordMessage("All this work just to do this?")
    who := runenv.TestInstanceParams["who"]
    runenv.RecordMessage("Hello, %s", who)
    
    return nil
}
```
{% endcode %}

### Execute the plan

Now comes the fun part -- see your test plan compiled and executed!

Open two terminals, one for the server and one for the client.

```go
$ testground daemon
```

```go
$ testground run single \
      --plan=quickstart \
      --testcase=testcase1 \
      --builder=docker:go \
      --runner=local:docker \
      --instances=1
      --collect
```

This will start a flurry of activity that will leave you wondering "what gives? isn't this a simple little hello world program?" Well, testground provides a few features that aren't  exercised by this example. Continue on with this tutorial to learn more about writing plans. In the mean time, here is a list of what you have just witnessed:

1. An isolated `testground-build` network is created
2. `goproxy` docker image is downloaded and a container is started
3. Your code is copied into a new docker image
4. Your code is compiled, and added to a new docker image
5. Several supporting containers are downloaded and started
   1. `testground-sidecar`
   2. `testground-redis` - synchronization service backend
   3. metrics \(Prometheus\)
   4. dashboards \(Grafana\)
6. A number of containers are created based on your image \(`--instances` controls this\)
7. Your code is executed in each of the containers
8. The `outputs` of the plan run are collected for analysis

