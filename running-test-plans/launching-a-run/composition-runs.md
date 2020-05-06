# Composition runs

A _"composition"_ run, is one in which multiple versions of the same software can be tested simultaneously. 

To define a composition run, a manifest for the _single_ plan must already exist. Indeed, a composition run is just multiple single runs being executed simultaneously with different versions of code being imported.

Here is how we go about creating a composition test run. The following example is taken from the libp2p test plans. Lets have a look at [this one](https://github.com/libp2p/test-plans/blob/master/dht/compositions/find-peers.example.toml) for an example, which is copied below.

#### Example composition

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

* code is divided into two groups. These goups can have any name. In this example, the groups are called **bootstrappers** and **peers**. 
* Both groups will execute the same plan, which is defined in the global section of this file.
* dependencies can be explicitly versioned, as indicated by the build build parameters of the boogstrappers group.

#### Building the composition

Before we can execute this example, make sure the plans are imported

```text
testground plan import --git --source https://github.com/libp2p/test-plans
```

With this plan imported \(including this composition file, we can now build the composition. Note that this will create multiple docker images if using the docker runner testground run 

```text
testground build -f <composition_toml> -w
```

#### Executing the composition

```text
testground run -f <composition_toml>
```



