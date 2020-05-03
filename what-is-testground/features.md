# Features

## 💡 Supports \(or aims to support\) a variety of testing workflows

> \(🌕 = fully supported // 🌑 = planned\)

* Experimental/iterative development 🌖 \(The team at Protocol Labs has used Testground extensively to evaluate protocol changes in large networks, simulate attacks, measure algorithmic improvements across network boundaries, etc.\)
* Debugging 🌗
* Comparative testing 🌖
* Backwards/forward-compatibility testing 🌖
* Interoperability testing 🌑
* Continuous integration 🌑
* Stakeholder/acceptance testing 🌑

## 📄 Simple, normalized, formal runtime environment for tests

A test plan is a blackbox with a formal contract. Testground promises to inject a set of env variables, and the test plan promises to emit events on stdout, and assets on the output directory.

* As such, a test plan can be any kind of program, written in Go, JavaScript, C, or shell.
* At present, we offer builders for Go, with TypeScript \(node and browser\) being in the works.

## 🛠 Modular builders and runners

For running test plans written in different languages, targeted for different runtimes, and levels of scale:

* `exec:go` and `docker:go` builders: compile test plans written in Go into executables or containers.
* `local:exec`, `local:docker`, `cluster:k8s` runners: run executables or containers locally \(suitable for 2-300 instances\), or in a Kubernetes cloud environment \(300-10k instances\).

> Got some spare cycles and would like to add support for writing test plans Rust, Python or X? It's easy! Open an issue, and the community will guide you!

## 👯‍♀️ Distributed coordination API

Redis-backed lightweight API offering synchronisation primitives to coordinate and choreograph distributed test workloads across a fleet of nodes.

## ☎️ Network traffic shaping

Test instances are able to set connectedness, latency, jitter, bandwidth, duplication, packet corruption, etc. to simulate a variety of network conditions.

## ☁️ Quickstart k8s cluster setup on AWS

Create a k8s cluster ready to run Testground jobs on AWS by following the instructions at [`testground/infra`](https://github.com/testground/infra).

## 🧩 Upstream dependency selection

Compiling test plans against specific versions of upstream dependencies \(e.g. moduleX v0.3, or commit 1a2b3c\).

## 🌱 Dealing with upstream API changes

So that a single test plan can work with a range of versions of the components under test, as these evolve over time.

## 📈 Results and diagnostics, raw and aggregated data points

**Diagnostics:** Automatic diagnostics via pprof \(for Go test plans\), with metrics emitted to InfluxDB in real-time. Metrics can be raw data points or aggregated measurements, such as histograms, counters, gauges, moving averages, etc.

**Results:** When the test plan concludes, all results are pushed in batch to InfluxDB for later exploration, analysis, and visualization.

## 🎼 Declarative jobs, we call them _compositions_

Create tailored test runs by composing scenarios declaratively, with different groups, cohorts, upstream deps, test params, etc.

## 💾 Emit and collect test outputs

Emit and collect/export/download test outputs \(logs, assets, event trails, run events, etc.\) from all participants in a run.

