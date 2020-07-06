# Running test plans

## Picking a runner

Plans which run on one runner generally should be usable on all other runners as well. The following table describes the features of different runners.‌

It is common practice when developing a test plan to use a local runner \(`local:exec` or `local:docker`\) in order to iterate quickly and then move to the Kubernetes `cluster:k8s` runner when you want to run your test plan with many more test instances.

|  runner | quick iteration | high instance count | network containment | traffic shaping | quick setup |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **local:exec** | ✅ | ❌ | ❌ | ❌ | ✅ |
| **local:docker** | ✅ | ❌ | ✅ | ✅ | ✅ |
| **cluster:k8s** | ❌ | ✅ | ✅ | ✅ | ❌ |

## Picking a builder

| builder | local:exec | local:docker | cluster:k8s | dependency mgmt | customizable |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **exec:go** | ✅ | ❌ | ❌ | ✅ | ❌ |
| **docker:go** | ❌ | ✅ | ✅ | ✅ | ❌ |
| **docker:generic** | ❌ | ✅ | ✅ | ❌ | ✅ |

## Single runs

A `single` run is one in which a single version of the code is being tested. This is the simplest way to run a test plan. If you have run through the [Writing test plans -&gt; Quick start](writing-test-plans/quickstart.md) tutorial, you have already encountered this:

```text
$ testground run single -p <plan> -t <testcase> -r <runner> -b <builder> -i <instances>
```

A single test run can be used to benchmark, validate, or observe the behaviour of a peer-to-peer system.

## Composition runs

A `composition` run, is one in which multiple versions of the same software can be tested simultaneously.  Where single runs allow you the performance of a single version of code, compositions allow you to test the way new developments interact with the existing network. This feature has been used heavily during libp2p DHT development.

To define a composition run, a manifest for the __`single` plan must already exist. Indeed, a composition run is just multiple single runs being executed simultaneously with different versions of code being imported.

Here is how we go about creating a composition test run. The following is a simple example. For a more complete example used by the libp2p project, I recommend you have a look at [libp2p compositions](https://github.com/libp2p/test-plans/tree/master/dht/compositions).

### Example composition

{% code title="composition.toml" %}
```text
[metadata]
name    = "quickstart"
author  = "your name here"

[global]
plan    = "mycomposition"
case    = "quickstart"
builder = "docker:go"
runner  = "local:docker"

total_instances = 5

[[groups]]
id = "group1"
instances = { count = 2 }

  [groups.build]
  selectors = ["foo"]
  dependencies = [
 #     { module = "github.com/your/module", version = "995fee9e5345fdd7c151a5fe871252262db4e788"},
  ]

  [groups.run]
  test_params = { }

[[groups]]
id = "group2"
instances = { count = 3 }

  [groups.build]

  [groups.run]
  test_params = { }
```
{% endcode %}

* Code is divided into two groups. These groups can have any name. In this example, the groups are called **group1** and **group2**. 
* Both groups will execute the same plan, which is defined in the global section of this file.
* Dependencies can be explicitly versioned. To force **group1** to use a specific software version, override the dependencies.

### Building a composition

Before we can execute this example, make sure we have a test plan we want to run. Create a new plan, or import an existing plan.

```text
$ testground plan create -p mycomposition
```

Copy the file above and into a file called `composition.toml`. Because of the way this file is written, plans will use the `local:docker` runner.

{% hint style="info" %}
Be aware that the composition file will be overwritten to include the names of built artifacts!
{% endhint %}

```bash
# build the artifacts and update the composition file
$ testground build composition -f composition.toml -w

# check the difference
$ cat composition.toml
```

### Executing a composition

```text
$ testground run composition -f composition.toml
```



