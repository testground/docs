---
description: 'How to install Testground, and run your first test plan'
---

# Getting started

## Installing Testground

### Prerequisites

* [Docker](https://www.docker.com/products/docker-desktop)
* [Go 1.16](https://golang.org/), or higher

### Installation

Currently, we don't distribute binaries, so you will have to build from source.

```bash
$ git clone https://github.com/testground/testground.git

$ cd testground

# compile Testground and all related dependencies
$ make install
```

## Running Testground

In order to use Testground, you need to have a running Testground daemon.

```bash
# start the Testground daemon, listening by default on localhost:8042
# it processes commands received from the Testground client
$ testground daemon
```

**`$TESTGROUND_HOME`** is an important directory. If not explicitly set, Testground uses `$HOME/testground` as a default. The layout of **`$TESTGROUND_HOME`** is as follows:

```text
$TESTGROUND_HOME
 |
 |__ plans              >>> [c] contains test plans, can be git checkouts, symlinks to local dirs, or the source itself
 |    |__ suite-a       >>> test plans can be grouped in suites (which in turn can be nested); this enables you to host many test plans in a single repo / directory.
 |    |    |__ plan-1   >>> source of a test plan identified by suite-a/plan-1 (relative to $TESTGROUND_HOME/plans) 
 |    |    |__ plan-2
 |    |__ plan-3        >>> source of a test plan identified by plan-3 (relative to $TESTGROUND_HOME/plans)
 |
 |__ sdks               >>> [c] hosts the test development SDKs that the client knows about, so they can be used with the --link-sdk option.
 |    |__ sdk-go
 |
 |__ data               >>> [d] data directory  
      |__ outputs
      |__ work

[c] = used client-side // [d] = used mostly daemon-side.
```

## Running an example test plan

The first test plan that we will run is the `network` test plan and the `ping-pong` test case.

The `ping-pong` test case starts 2 test plan instances: one that listens on a TCP socket and another that dials it. The test case exercises the [sync service](concepts-and-architecture/sync-service.md) as well as the [traffic shaping](traffic-shaping.md) and IP allocation functionality.

Configure `$TESTGROUND_HOME` and copy the example `network` test plan into the `$TESTGROUND_HOME/plans` directory.

```bash
# assuming you already started your Testground daemon (as instructed above)
# there should be a `testground` directory in your home folder, i.e. `~/testground`
#
# from your testground/testground Git checkout, run:
$ testground plan import --from ./plans/network
...
created symlink /Users/raul/testground/plans/network -> ./plans/network
imported plans:
network ping-pong
```

Run the `network`testplan and the `ping-pong` test case with the `docker:go` builder and the `local:docker` runner.

{% hint style="info" %}
Make sure you have `testground daemon` running in another terminal window.
{% endhint %}

```bash
$ testground run single \
         --plan=network \
         --testcase=ping-pong \
         --builder=docker:go \
         --runner=local:docker \
         --instances=2 \
         --wait
```

{% hint style="info" %}
During the first run the Testground daemon sets up the builder and runner environments. Subsequent runs will be faster.
{% endhint %}

You should see a flurry of activity, including measurements, messages, and runtime events. When the execution concludes, you will see something like:

```text
[...]
INFO run finished successfully {"req_id": "d570c53a", "plan": "network", "case": "ping-pong", "runner": "local:docker", "instances": 2}

>>> Result:

INFO finished run with ID: 5222e5df793b
```

In the local runners, all test plan run outputs and logs are stored at `$TESTGROUND_HOME/data`. Collect them into a bundle with the following command \(replacing `5222e5df793b` with the corresponding run ID\):

```bash
$ testground collect --runner=local:docker 5222e5df793b
[...]

>>> Result:

INFO    created file: 5222e5df793b.tgz
```

Open the bundle and you will find the outputs from the test in there:

![](.gitbook/assets/image%20%282%29%20%282%29.png)

## Configuration \(.env.toml\)

`.env.toml`is a configuration file read by the Testground daemon and the Testground client on startup.

Testground tries to load this file from `$TESTGROUND_HOME/.env.toml`, where `$TESTGROUND_HOME` defaults to `$HOME/testground` by default.

### Changing default daemon bind addresses

You can change the default bind addresses by configuring `daemon.listen` and `client.endpoint`

{% code title=".env.toml" %}
```text
[daemon]
listen = ":8080"

[client]
endpoint = "http://localhost:8080"
```
{% endcode %}

The endpoint refers to the `testground-daemon` service, so depending on your setup, this could be, for example, a Load Balancer fronting the kubernetes cluster and forwarding proper requests to the `tg-daemon` service, or a simple port forward to your local workstation:

```
[client]
endpoint = "http://localhost:28015" # in case we use port forwarding, like this one here: kubectl port-forward service/testground-daemon 28015:8042
```

### Customize asynchrony

You can customize the number of asynchronous workers, as well as the maximum queue capacity, i.e., the maximum number of pending tasks at a moment in time.

{% code title=".env.toml" %}
```text
[daemon]
workers = 2
queue_size = 100
```
{% endcode %}

### AWS integration

When using a remote runner such as `cluster:k8s`, you should configure the default region:

{% code title=".env.toml" %}
```text
["aws"]
region = "aws region, such as eu-central-1"
```
{% endcode %}

The AWS configuration is also used if you push Docker images to AWS ECR from the `docker:go` builder using the `--build-cfg push_registry=true` and `--build-cfg registry_type=aws` flags.

### DockerHub integration

If you want to push Docker images from the `docker:go` builder to a DockerHub registry, you can configure it.

{% code title=".env.toml" %}
```text
["dockerhub"]
repo = "repo to be used for testground"
username = "username"
access_token = "docker hub access token"
```
{% endcode %}

