# Getting started

## Installing Testground

### Prerequisites

* [Docker](https://www.docker.com/products/docker-desktop)
* [Go 1.14](https://golang.org/) or higher

### Installation

Currently, we don't distribute binaries, so you will have to build from source.

```bash
$ git clone https://github.com/testground/testground.git

$ cd testground

# compile Testground and all related dependencies
$ make install

# start the Testground daemon, it listens on localhost:8042 by default
$ testground daemon
```

**`$TESTGROUND_HOME`** is an important directory. If not explicitly set, Testground uses `$HOME/testground` as a default.

The layout of **`$TESTGROUND_HOME`** is as follows:

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

The `ping-pong` test case starts 2 test plan instances: one that listens on a TCP socket and another that dials it. The test case exercises the synchronization service as well as the sidecar traffic shaping and IP allocation functionality.

Configure `$TESTGROUND_HOME` to a directory that would hold your test plan:

```
$ mkdir -p ~/.testground

$ export TESTGROUND_HOME=~/testground
```

Copy the `network` test plan into the `$TESTGROUND_HOME` directory.

```
$ mkdir -p $TESTGROUND_HOME/plans

$ cd $GOPATH/src/github.com/testground/testground

$ cp -r plans/network $TESTGROUND_HOME/plans
```

Run the `network`testplan and the `ping-pong` test case with the `docker:go` builder and the `local:docker` runner.

```
$ testground run single \
                 --plan=network \
                 --testcase=ping-pong \
                 --builder=docker:go \
                 --runner=local:docker \
                 --instances=2
```

After the test plan run execution concludes, you should see the following message:

```
INFO run finished successfully {"req_id": "d570c53a", "plan": "network", "case": "ping-pong", "runner": "local:docker", "instances": 2}

INFO finished run with ID: 5222e5df793b
```

All test plan run outputs and logs are stored at `$TESTGROUND_HOME/data/outputs/local_docker/network`

You could also fetch them with the following command:

```
$ testground collect --runner=local:docker 5222e5df793b
```

## Configuration \(.env.toml\)

`.env.toml`is a configuration file read by the Testground daemon and Testground client on startup.

Testground tries to load this file from `$TESTGROUND_HOME/.env.toml`

### Changing default bind addresses

You can change the default bind addresses by configuring `daemon.listen` and `client.endpoint`

{% code title=".env.toml" %}
```text
[daemon]
listen = ":8080"

[client]
endpoint = "localhost:8080"
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

### DockerHub integration

If you want to push Docker images from the `docker:go` builder to a remote registry, you can configure it.

{% code title=".env.toml" %}
```text
["dockerhub"]
repo = "repo to be used for testground"
username = "username"
access_token = "docker hub access token"
```
{% endcode %}



