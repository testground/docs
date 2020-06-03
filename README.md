# What is Testground?

![](.gitbook/assets/image%20%283%29.png)

{% hint style="info" %}
You are reading the Testground documentation for the `master` branch.

The Testground team maintains documentation for the `master` branch and for the latest stable release.
{% endhint %}

## Overview

Testground is a platform for testing, benchmarking, and simulating distributed and peer-to-peer systems at scale. It's designed to be multi-lingual and runtime-agnostic, scaling gracefully from 2 to 10k instances, only when needed.

The Testground project was started at Protocol Labs because we couldn't find a platform that would allow us to reliably and reproducibly test and measure how changes to the IPFS and libp2p codebases would impact the performance and health of large networks \(as well as individual nodes\), so we decided to invent it.

![](.gitbook/assets/testground-demo.gif)

## How does it work?

### 1. **You develop distributed test plans as if you were writing unit tests against local APIs**

* No puppeteering necessary.
* No need to package and ship the system as a separate daemon with an external API in order to puppeteer it.
* No need to expose every internal setting over an external API, just for the sake of testing.

### **2. Your test plan calls out to the coordination API to...**

* communicate out-of-band information \(such as endpoint addresses, peer ids, etc.\)
* leverage synchronization and ordering primitives such as signals and barriers to model a distributed state machine.
* programmatically apply network traffic shaping policies, which you can alter during the execution of a test to simulate various network conditions.

### **3. There is no special "conductor" node telling instances what to do when**

* The choreography and sequencing emerges from within the test plan itself.

### **4. You decide what versions of the upstream software you want to exercise your test against.**

* Benchmark, simulate, experiment, run attacks, etc. against versions v1.1 and v1.2 of the components under test in order to compare results, or test compatibility.
* Assemble hybrid test runs mixing various versions of the dependency graph.

### **5. Inside your test plan...**

* You record observations, metrics, success/failure statuses.
* You emit structured or unstructured assets you want collected, such as event logs, dumps, snapshots, binary files, etc.

### **6. Via a TOML-based** _**composition**_ **file, you instruct Testground to...**

* Assemble a test run comprising groups of 2, 200, or 10000 test instances, each with different test parameters, or built against different dependency sets.
* Schedule them for run locally \(executable or Docker\), or in a cluster \(Kubernetes\).

### **7. You collect the outputs of the test plan for analysis...**

* with a single command...
* using data processing scripts and platforms \(such as the upcoming Jupyter notebooks integration\) to draw conclusions.

## Features

### 💡 Supports \(or aims to support\) a variety of testing workflows

> \(🌕 = fully supported // 🌑 = planned\)

* Experimental/iterative development 🌖
  * The team at Protocol Labs has used Testground extensively to evaluate protocol changes in large networks, simulate attacks, measure algorithmic improvements across network boundaries, etc.
* Debugging 🌗
* Comparative testing 🌖
* Backwards/forward-compatibility testing 🌖
* Interoperability testing 🌑
* Continuous integration 🌑
* Stakeholder/acceptance testing 🌑

### 📄 Simple, normalized, formal runtime environment for tests

A test plan is a blackbox with a formal contract. Testground promises to inject a set of env variables, and the test plan promises to emit events on stdout, and assets on the output directory.

* As such, a test plan can be any kind of program, written in Go, JavaScript, C, or shell.
* At present, we offer builders for Go, with TypeScript \(node and browser\) being in the works.

### 🛠 Modular builders and runners

For running test plans written in different languages, targeted for different runtimes, and levels of scale:

* `exec:go` and `docker:go` builders: compile test plans written in Go into executables or containers.
* `local:exec`, `local:docker`, `cluster:k8s` runners: run executables or containers locally \(suitable for 2-300 instances\), or in a Kubernetes cloud environment \(300-10k instances\).

> Got some spare cycles and would like to add support for writing test plans Rust, Python or X? It's easy! Open an issue, and the community will guide you!

### 👯‍♀️ Distributed coordination API

Redis-backed lightweight API offering synchronisation primitives to coordinate and choreograph distributed test workloads across a fleet of nodes.

### ☎️ Network traffic shaping

Test instances are able to set connectedness, latency, jitter, bandwidth, duplication, packet corruption, etc. to simulate a variety of network conditions.

### ☁️ Quickstart k8s cluster setup on AWS

Create a k8s cluster ready to run Testground jobs on AWS by following the instructions at [`testground/infra`](https://github.com/testground/infra).

### 🧩 Upstream dependency selection

Compiling test plans against specific versions of upstream dependencies \(e.g. moduleX v0.3, or commit 1a2b3c\).

### 🌱 Dealing with upstream API changes

So that a single test plan can work with a range of versions of the components under test, as these evolve over time.

### 📈 Results and diagnostics, raw and aggregated data points

**Diagnostics:** Automatic diagnostics via pprof \(for Go test plans\), with metrics emitted to InfluxDB in real-time. Metrics can be raw data points or aggregated measurements, such as histograms, counters, gauges, moving averages, etc.

**Results:** When the test plan concludes, all results are pushed in batch to InfluxDB for later exploration, analysis, and visualization.

### 🎼 Declarative jobs, we call them _compositions_

Create tailored test runs by composing scenarios declaratively, with different groups, cohorts, upstream deps, test params, etc.

### 💾 Emit and collect test outputs

Emit and collect/export/download test outputs \(logs, assets, event trails, run events, etc.\) from all participants in a run.

## Documentation issues

{% hint style="warning" %}
**This docs site is work-in-progress!** You're bound to find dragons 🐉 in some sections, so please bear with us! If something looks wrong, please [open a docs issue](https://github.com/testground/testground/issues/new?assignees=&labels=docs&template=DOCS.md&title=docs%20site:%20%3Cdescribe%20the%20problem%3E) on our main repo.
{% endhint %}

