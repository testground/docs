# Running test plans

## Picking a runner

Plans which run on one runner generally should run be runnable on all other runners as well. The following table describes the features of different runners.‌

It is common practice when developing a test plan to use a local runner \(`local:exec` or `local:docker`\) in order to iterate quickly and then move to the Kubernetes `cluster:k8s` runner when you want to run your test plan with many more test instances.

|  runner | quick iteration | high instance count | network containment | quick setup |
| :--- | :--- | :--- | :--- | :--- |
| **local:exec** | ✅ | ❌ | ❌ | ✅ |
| **local:docker** | ✅ | ❌ | ✅ | ✅ |
| **cluster:k8s** | ❌ | ✅ | ✅ | ❌ |

## Single runs

A `single` run, is one in which a single version of the code is being tested. This is the simplest way to run a test plan. If you have run through the [Writing test plans -&gt; Quick start](writing-test-plans/quickstart.md) tutorial, you have already encountered this:

```text
$ testground run single -p <plan> -t <testcase> -r <runner> -b <builder> -i <instances>
```

A single test run can be used to benchmark, validate, or observe the behaviour of a peer-to-peer system.

## Composition runs

A `composition` run, is one in which multiple versions of the same software can be tested simultaneously. 

To define a composition run, a manifest for the __`single` plan must already exist. Indeed, a composition run is just multiple single runs being executed simultaneously with different versions of code being imported.

Here is how we go about creating a composition test run. The following example is taken from the libp2p test plans. Let's have a look at [this one](https://github.com/libp2p/test-plans/blob/master/dht/compositions/find-peers.example.toml) for an example, which is copied below.

### Example composition

{% code title="find-peers.toml" %}
```text
[metadata]
name    = "find-peers-01"
author  = "raulk"

[global]
plan    = "dht"
case    = "find-peers"
builder = "exec:go"
runner  = "local:exec"

total_instances = 50

[[groups]]
id = "bootstrappers"
instances = { count = 1 }

  [groups.build]
  selectors = ["foo"]
  dependencies = [
      { module = "github.com/libp2p/go-libp2p-kad-dht", version = "995fee9e5345fdd7c151a5fe871252262db4e788"},
      { module = "github.com/libp2p/go-libp2p", version = "76944c4fc848530530f6be36fb22b70431ca506c"},
  ]

  [groups.run]
  test_params = { random_walk = "true", n_bootstrap = "1" }

[[groups]]
id = "peers"
instances = { count = 49 }

  [groups.run]
  test_params = { random_walk = "true", n_bootstrap = "1" }

  [groups.build]

```
{% endcode %}

* Code is divided into two groups. These groups can have any name. In this example, the groups are called **bootstrappers** and **peers**. 
* Both groups will execute the same plan, which is defined in the global section of this file.
* Dependencies can be explicitly versioned, as indicated by the build parameters of the **bootstrappers** group.

### Building a composition

Before we can execute this example, make sure the plans are imported

```text
$ testground plan import --git --source https://github.com/libp2p/test-plans
```

With the plan imported \(including this composition file\), we can now build the composition. Note that this will create multiple Docker images if you're using the `local:docker` runner.

```text
$ testground build -f find-peers.toml -w
```

### Executing a composition

```text
$ testground run -f find-peers.toml
```



