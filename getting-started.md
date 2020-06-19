---
description: 'How to install Testground, and run your first test plan'
---

# Getting started

## Installing Testground

### Prerequisites

* [Docker](https://www.docker.com/products/docker-desktop)
* [Go 1.14](https://golang.org/), or higher

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
         --instances=2
```

{% hint style="info" %}
During the first run the Testground daemon sets up the builder and runner environments. Subsequent runs will be faster.
{% endhint %}

You should see a flurry of activity, including measurements, messages, and runtime events. When the execution concludes, you will see something like:

```
[...]
INFO run finished successfully {"req_id": "d570c53a", "plan": "network", "case": "ping-pong", "runner": "local:docker", "instances": 2}

>>> Result:

INFO finished run with ID: 5222e5df793b
```

In the local runners, all test plan run outputs and logs are stored at `$TESTGROUND_HOME/data`.  Collect them into a bundle with the following command \(replacing `5222e5df793b` with the corresponding run ID\):

```bash
$ testground collect --runner=local:docker 5222e5df793b
[...]

>>> Result:

INFO	created file: 5222e5df793b.tgz
```

Open the bundle and you will find the outputs from the test in there:

![](.gitbook/assets/image%20%281%29.png)

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

## Resolving issues with the Windows Subsystem for Linux

Version 2 of Microsoft's Windows Subsystem for Linux \(WSL2\) enables a new backend for Docker Desktop in which Docker runs on the same virtual machine as the Linux userland instead of in a separate VM. This makes Docker much faster and uses system resource more efficiently, but the default Linux kernel for WSL2 is not compatible with testground's [traffic shaping](traffic-shaping.md) feature.

Note that if you're using version 1 of WSL, or if Docker Desktop is not configured to use the WSL2 backend, you should not be affected by this issue. Affected users will see errors similar to the following when running a test plan that uses traffic shaping:

```text
ERROR   sidecar worker failed: failed to initialise the container: failed to initialize link default (eth1): failed to set root qdisc: no such file or directory
```

To enable traffic shaping under WSL2, you'll need to build a Linux kernel with a custom configuration.

First, install the requirements for building the kernel:

```text
sudo apt install build-essential flex bison libssl-dev libelf-dev
```

This assumes you're using a debian based distribution like Ubuntu; adapt to your preferred package manager otherwise.

Then clone the kernel sources:

```text
git clone --depth=1 https://github.com/microsoft/WSL2-Linux-Kernel.git
cd WSL2-Linux-Kernel
```

And copy Microsoft's kernel config to use as a base:

```text
cp Microsoft/config-wsl .config
```

Now edit the `.config` file and find the `Queueing/Scheduling` section. We want it to look like this:

```text
#
# Queueing/Scheduling
#
CONFIG_NET_SCH_CBQ=y
CONFIG_NET_SCH_HTB=y
CONFIG_NET_SCH_HFSC=y
CONFIG_NET_SCH_PRIO=y
CONFIG_NET_SCH_MULTIQ=y
CONFIG_NET_SCH_RED=y
CONFIG_NET_SCH_SFB=y
CONFIG_NET_SCH_SFQ=y
CONFIG_NET_SCH_TEQL=y
CONFIG_NET_SCH_TBF=y
CONFIG_NET_SCH_CBS=y
CONFIG_NET_SCH_ETF=y
CONFIG_NET_SCH_GRED=y
CONFIG_NET_SCH_DSMARK=y
CONFIG_NET_SCH_NETEM=y
CONFIG_NET_SCH_DRR=y
CONFIG_NET_SCH_MQPRIO=y
CONFIG_NET_SCH_SKBPRIO=y
CONFIG_NET_SCH_CHOKE=y
CONFIG_NET_SCH_QFQ=y
CONFIG_NET_SCH_CODEL=y
CONFIG_NET_SCH_FQ_CODEL=y
CONFIG_NET_SCH_CAKE=y
CONFIG_NET_SCH_FQ=y
CONFIG_NET_SCH_HHF=y
CONFIG_NET_SCH_PIE=y
CONFIG_NET_SCH_INGRESS=y
CONFIG_NET_SCH_PLUG=y
CONFIG_NET_SCH_DEFAULT=y
# CONFIG_DEFAULT_FQ is not set
# CONFIG_DEFAULT_CODEL is not set
CONFIG_DEFAULT_FQ_CODEL=y
# CONFIG_DEFAULT_SFQ is not set
# CONFIG_DEFAULT_PFIFO_FAST is not set
CONFIG_DEFAULT_NET_SCH="fq_codel"
```

It's also a good idea to change the `CONFIG_LOCALVERSION` setting from `"-microsoft-standard"` to something recognizable, e.g. `"-wsl-traffic-shaping"`.

Now you can build the kernel:

```text
make
```

When it's finished, copy the compiled kernel somewhere on your Windows filesystem. For this example, we'll put the kernel at `C:\wsl\kernel-traffic-shaping.bzImage`.

Create the `c:\wsl` folder if needed, then copy the file:

```text
cp arch/x86_64/boot/bzImage /mnt/c/wsl/kernel-traffic-shaping.bzImage
```

To use the new kernel, edit the `.wslconfig` file in your Windows home directory, creating it if it doesn't exist:

```text
[wsl2]
kernel=C:\\wsl\\kernel-traffic-shaping.bzImage
```

Notice that you need to use Windows-style paths with escaped backslashes.

Now from a Windows command prompt, run `wsl.exe --shutdown` - this will halt the Linux VM, so make sure you don't have any unsaved work in Linux land.

Opening a new WSL session should use the new kernel. You can verify this with `uname -a`, which should show the value you set for `CONFIG_LOCALVERSION` earlier instead of `-microsoft-standard`.

Now, testground traffic shaping should work as expected. You can verify this by running the `network/ping-pong` plan included in the testground repository.



