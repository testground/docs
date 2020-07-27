# Table of contents

* [What is Testground?](README.md)
* [Concepts and architecture](concepts-and-architecture/README.md)
  * [Test plans and test cases](concepts-and-architecture/test-structure.md)
  * [Daemon and client](concepts-and-architecture/daemon-and-client.md)
  * [Synchronization service](concepts-and-architecture/sync-service.md)
  * [Networking](concepts-and-architecture/networking.md)
  * [Sidecar](concepts-and-architecture/sidecar.md)
  * [Builders](concepts-and-architecture/builders.md)
  * [Runners](concepts-and-architecture/runners.md)
  * [Runtime environment \(runenv\)](concepts-and-architecture/runtime.md)
  * [Client-Server communication](concepts-and-architecture/client-server-communication.md)
* [Getting started](getting-started.md)
* [Writing test plans](writing-test-plans/README.md)
  * [Quick start](writing-test-plans/quickstart.md)
  * [Understanding the test plan manifest](writing-test-plans/test-plan-manifest.md)
  * [Parameters and test cases](writing-test-plans/paramaters-and-testcases.md)
  * [Keeping instances in sync](writing-test-plans/synchronization.md)
  * [Communication between instances](writing-test-plans/communication-between-instances.md)
  * [Observability, assets and metrics](writing-test-plans/observability-assets-and-metrics.md)
* [Managing test plans](managing-test-plans.md)
* [Running test plans](running-test-plans.md)
* [Traffic shaping](traffic-shaping.md)
* [Analyzing test run results](analyzing-the-results.md)
* [Debugging test plans](debugging-test-plans.md)

## Runner library

* [local:exec](runner-library/local-exec.md)
* [local:docker](runner-library/local-docker/README.md)
  * [System overview](runner-library/local-docker/system-overview.md)
  * [Runner flags](runner-library/local-docker/runner-flags.md)
  * [Troubleshooting](runner-library/local-docker/troubleshooting.md)
* [cluster:k8s](runner-library/cluster-k8s/README.md)
  * [System overview](runner-library/cluster-k8s/system-overview.md)
  * [How to create a Kubernetes cluster for Testground](runner-library/cluster-k8s/how-to-create-a-kubernetes-cluster-for-testground.md)
  * [Monitoring and Observability](runner-library/cluster-k8s/monitoring.md)
  * [Understanding Testground performance on Kubernetes](runner-library/cluster-k8s/understanding-testground-performance-on-kubernetes.md)
  * [Troubleshooting](runner-library/cluster-k8s/troubleshooting.md)

## BUILDER LIBRARY

* [docker:go](builder-library/docker-go.md)
* [exec:go](builder-library/exec-go.md)
* [docker:generic](builder-library/docker-generic.md)

