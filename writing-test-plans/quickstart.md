---
description: Writing your first test plan
---

# Quick start

In this quick start tutorial you will get up and running with a simple test plan. Later tutorials will go deeper into features of the Testground SDK and how to use it to simulate P2P network environments. For now, let's get your hands dirty writing your first test plan.

Following these steps you will create a test plan along with all supporting files.

Let's dive right in, and we'll explain as we go.

### 1. Install Testground

Go through the [Getting started](../getting-started.md) page and follow the instructions in order to install Testground.

### 2. Create a test plan

Testground stores plans in the `$TESTGROUND_HOME` directory. This directory is created for you in your `home` directory, but if you prefer to store plans in another location, you can adjust the location using the environment variable.

```bash
# OPTIONAL - adjust TESTGROUND_HOME, default is ~/testground
$ export TESTGROUND_HOME=/path/to/testground/home

# create your first test plan
$ testground plan create --plan=quickstart
[...]
test plan created under: $TESTGROUND_HOME/plans/quickstart

# Bingo! A new, barebones templated Go plan will have been created
# under $TESTGROUND_HOME/plans/quickstart; explore the source!
```

### 3. Start the Testground daemon

```text
$ testground daemon
```

### 4. Run the plan

```bash
$ testground run single --plan=quickstart \
                        --testcase=quickstart \
                        --runner=local:docker \
                        --builder=docker:go \
                        --instances=1 \
                        --wait
```

### You did it!

This will start a flurry of activity that will leave you wondering _"what gives? isn't this a simple little hello world program?"_.

Well, Testground provides a few features that aren't exercised by this example. Continue with this tutorial to learn more about writing test plans.

Here is a list of what you have just witnessed:

1. An isolated `testground-build` network is created.
2. `goproxy` Docker image is downloaded and a container is started.
3. Your code is copied into a new Docker image.
4. Your code is compiled, and added to a new Docker image.
5. Several supporting containers are downloaded and started:
   1. `testground-sidecar` - the Testground service responsible for adjusting the test plan networks.
   2. `testground-redis` - the Testground synchronization service backend.
   3. `influxdb` - runtime metrics from test plans.
   4. `grafana` - dashboards for monitoring test plans and infrastructure.
6. A number of containers are created based on your image \(`--instances` controls this\).
7. Your code is executed in each of the containers.
8. The `outputs` of the test run are collected for analysis.

