# Table of contents

* [What is Testground?](README.md)
* [Concepts and architecture](concepts-and-architecture/README.md)
  * [Test plans and test cases](concepts-and-architecture/test-structure.md)
  * [Daemon and client](concepts-and-architecture/daemon-and-client.md)
  * [Synchronization service](concepts-and-architecture/sync-service.md)
  * [Networking](concepts-and-architecture/networking.md)
  * [Sidecar](concepts-and-architecture/sidecar.md)
  * [Runners](concepts-and-architecture/runners.md)
  * [Builders](concepts-and-architecture/builders-1.md)
  * [Environment variables](concepts-and-architecture/runtime.md)
  * [Runtime environment \(runenv\)](concepts-and-architecture/runtime-environment-runenv.md)
* [Getting started](getting-started.md)
* [Writing test plans](writing-test-plans/README.md)
  * [Test plan manifest](writing-test-plans/test-plan-manifest.md)
  * [Shimming](writing-test-plans/shimming.md)
* [Managing test plans](managing-test-plans.md)
* [Running test plans](running-test-plans/README.md)
  * [Picking a runner](running-test-plans/picking-a-runner.md)
  * [Creating instance groups](running-test-plans/creating-instance-groups.md)
  * [Launching a run](running-test-plans/launching-a-run/README.md)
    * [Composition runs](running-test-plans/launching-a-run/composition-runs.md)
    * [Single runs](running-test-plans/launching-a-run/single-runs.md)
  * [Setting test parameters](running-test-plans/test-parameters.md)
* [Traffic shaping](traffic-shaping.md)
* [Analyzing the results](analyzing-the-results.md)
* [Debugging test plans](debugging-test-plans.md)

## local:exec runner

* [System overview](local-exec-runner/system-overview.md)

## local:docker runner

* [System overview](local-docker-runner/system-overview.md)
* [Runner flags](local-docker-runner/runner-flags.md)
* [Monitoring](local-docker-runner/monitoring.md)
* [Troubleshooting](local-docker-runner/troubleshooting.md)

## cluster:k8s runner

* [System overview](cluster-k8s-runner/system-overview.md)
* [How to create a Kubernetes cluster for Testground](cluster-k8s-runner/how-to-create-a-kubernetes-cluster-for-testground.md)
* [Monitoring and Observability](cluster-k8s-runner/monitoring.md)
* [Troubleshooting](cluster-k8s-runner/troubleshooting.md)
* [Understanding Testground performance on Kubernetes](cluster-k8s-runner/understanding-testground-performance-on-kubernetes.md)

## Test Plan SDK

* [Quickstart](test-plan-sdk/quickstart.md)
* [Parameters and test cases](test-plan-sdk/paramaters-and-testcases.md)
* [Keeping instances in sync](test-plan-sdk/synchronization.md)
* [Communication between instances](test-plan-sdk/communication-between-instances.md)

