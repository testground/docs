---
description: Writing your first test plan
---

# Quickstart

## Hello, Test Plans!

In this quick start tutorial you will get up and running with a simple test plan. Later tutorials will go deeper into features of the plan SDK and how to use it to simulate P2P network environments. But for now, let's get your hands dirty writing your first test plan. Hello!

Following this steps will create a test plan allong with all supporing files. Lets dive right in, and and I'll explain as we go.

### 1. Install testground \(if you haven't already\)

```bash
go get github.com/testground/testground
cd $GOPATH/src/github.com/testground/testground
make install
```

### 2. Create a plan

testground stores plans in a directory called the `TESTGROUND_HOME`. Typically this is created for you in your home directory, but if you prefer to store plans in another location, you can adjust the location using this environment variable.

```bash
export TESTGROUND_HOME=/path/to/testground/home #optional
testground plan create -p newplan
```

### 3. Start testground daemon

```text
testground daemon
```

### 4. Run the plan!

```text
testground run single -p newplan -t quickstart -r local:docker -b docker:go -i 1
```

### You did it!

This will start a flurry of activity that will leave you wondering "what gives? isn't this a simple little hello world program?" Well, testground provides a few features that aren't  exercised by this example. Continue on with this tutorial to learn more about writing plans. In the mean time, here is a list of what you have just witnessed:

1. An isolated `testground-build` network is created
2. `goproxy` docker image is downloaded and a container is started
3. Your code is copied into a new docker image
4. Your code is compiled, and added to a new docker image
5. Several supporting containers are downloaded and started
   1. `testground-sidecar`
   2. `testground-redis` - synchronization service backend
   3. metrics \(influxdb\)
   4. dashboards \(Grafana\)
6. A number of containers are created based on your image \(`--instances` controls this\)
7. Your code is executed in each of the containers
8. The `outputs` of the plan run are collected for analysis

