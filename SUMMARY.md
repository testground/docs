# Table of contents

* [What is Testground?](README.md)
* [Test plans and test cases](test-structure.md)
* [Concepts and architecture](concepts-and-architecture/README.md)
  * [Daemon and client](concepts-and-architecture/daemon-and-client.md)
  * [Synchronization service](concepts-and-architecture/sync-service.md)
  * [Networking](concepts-and-architecture/networking.md)
  * [Sidecar](concepts-and-architecture/sidecar.md)
  * [Runners and builders](concepts-and-architecture/runners-and-builders.md)
  * [Environment variables](concepts-and-architecture/runtime.md)
  * [Runtime environment \(runenv\)](concepts-and-architecture/runtime-environment-runenv.md)
* [Getting started](getting-started/README.md)
  * [Installing Testground](getting-started/installing-testground.md)
  * [Running an example test plan](getting-started/running-an-example.md)
  * [Configuration \(.env.toml\)](getting-started/configuration-env.toml.md)
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
* [What are runners?](runners-1.md)
* [Traffic shaping](traffic-shaping.md)
* [Analyzing the results](analyzing-the-results.md)
* [Debugging test plans](debugging-test-plans.md)
* [Monitoring](monitoring.md)
* [Troubleshooting](troubleshooting.md)

## Builders

* [What are Builders?](builders/builders.md)
* [exec:go builder](builders/exec-go-builder.md)
* [docker:go builder](builders/docker-go-builder.md)

## cluster:k8s runner

* [System overview](cluster-k8s-runner/system-overview.md)
* [How to create a Kubernetes cluster for Testground](cluster-k8s-runner/how-to-create-a-kubernetes-cluster-for-testground.md)
* [Troubleshooting](cluster-k8s-runner/troubleshooting.md)
* [Understanding Testground performance on Kubernetes](cluster-k8s-runner/understanding-testground-performance-on-kubernetes.md)

## Test Plan SDK

* [Quickstart](test-plan-sdk/quickstart.md)
* [Parameters and test cases](test-plan-sdk/paramaters-and-testcases.md)

## Sync service recipes <a id="sync-service-recipes-1"></a>

* [Keeping instances in sync](sync-service-recipes-1/synchronization.md)
* [Communication between instances](sync-service-recipes-1/communication-between-instances.md)

