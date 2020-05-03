# Table of contents

* [Welcome!](README.md)
* [What is Testground?](what-is-testground.md)
* [How tests are structured](test-structure/README.md)
  * [Test plans](test-structure/test-plans.md)
  * [Test cases](test-structure/test-cases.md)
  * [Test run](test-structure/test-run.md)
* [Concepts and architecture](concepts-and-architecture/README.md)
  * [Daemon and client](concepts-and-architecture/daemon-and-client.md)
  * [Sync service](concepts-and-architecture/sync-service.md)
  * [Sidecar](concepts-and-architecture/sidecar.md)
  * [Runners and builders](concepts-and-architecture/runners-and-builders.md)
  * [Runtime environment \(runenv\)](concepts-and-architecture/runtime-environment-runenv.md)
* [Getting started](getting-started/README.md)
  * [Prerequisites](getting-started/prerequisites.md)
  * [Installation and setup](getting-started/installation-and-setup.md)
  * [Running an example](getting-started/running-an-example.md)
  * [Configuration \(env.toml\)](getting-started/configuration-env.toml.md)
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
  * [Test parameters](running-test-plans/test-parameters.md)
* [Network simulation](network-simulation.md)
* [Analyzing the results](analyzing-the-results.md)
* [Debugging test plans](debugging-test-plans.md)
* [Sync service recipes](sync-service-recipes.md)
* [Monitoring](monitoring.md)
* [Troubleshooting](troubleshooting.md)

## Builder library

* [exec:go builder](builder-library/exec-go-builder.md)
* [docker:go builder](builder-library/docker-go-builder.md)

## Runner library

* [local:exec runner](runner-library/local-exec-runner.md)
* [local:docker runner](runner-library/local-docker-runner.md)
* [cluster:k8s runner](runner-library/cluster-k8s-runner.md)

## Executing plans

* [Builders](executing-plans/builders.md)
* [Runners](executing-plans/runners.md)

## Test Plan SDK

* [Quickstart](test-plan-sdk/quickstart.md)
* [Environment Variables](test-plan-sdk/runtime.md)
* [Paramaters and TestCases](test-plan-sdk/paramaters-and-testcases.md)
* [Keeping instances in sync](test-plan-sdk/synchronization.md)
* [Communication between instances](test-plan-sdk/communication-between-instances.md)
* [Testground Support Landing Page](test-plan-sdk/testground-support-landing-page-1.md)

