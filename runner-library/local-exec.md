# local:exec

## Background

Using the `local:exec` runner  enables you to test distributed/p2p software as quickly as possible. When using this runner, plans are compiled and executed as a standard process.

This runner is the only one which does not use a sidecar to modify the networking. Processes run un-contained. Any files created or network calls performed during the test will be visible on the host system.

## Dependencies

Although the plan binaries are compiled for the host system, all the auxiliary infrastructure is provided by Docker. This configuration enables easy installation of the sync service and monitoring infrastructure while maintaining quick build-execute iteration cycle which is the advantage of this runner.

* A laptop or desktop with reasonable hardware specs. If your computer is newer than 5 or 6 years old, it will probably be just fine. To be able run simulations of a reasonable size, we recommend at least the following:
  * 8GB memory
  * 50GB available storage
* The following software setup is required
  * docker daemon
  * a non-root account with write access to the docker socket.

## Testground-supplied containers

When Testground runs its first plan, several additional containers will be started. Here is an overview of everything Testground adds to your system

* Docker containers:
  * `redis`
    * This is the backend database of the Testground sync service.
  * `grafana`
    * Visualisation software that Testground uses for its dashboards.
  * `influxdb`
    * Time-series database that Testground uses for various diagnostics and events.

## Troubleshooting

Many of the infrastructure pieces provided by Testground are common with the `local:docker` runner, therefore the troubleshooting steps of the `local:docker` runner are applicable here as well.

