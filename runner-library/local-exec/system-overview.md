# System overview

## Background

Using the `local:exec` runner  enables you to test distributed/p2p software as quickly as possible. When using this runner, plans are compiled and executed as a standard process.

This runner is the only one which does not use a sidecar to modify the networking. Processes run un-contained. Any files created or network calls performed during the test will be visible on the host system.

## Dependencies:

Although the plan binaries are compiled for the host system, all the auxilary infrastructure is provided by docker. This configuration enables easy installation of the sync serice and monitoring infrastructure while maintinaing quick build-execute cycle which is the advantage of this runner.

* A laptop or desktop with reasonable hardware specs. If your computer is newer than 5 or 6 years old, it will probably be just fine. To be able run simulations of a reasonable size, we recommend at least the following:
  * 8GB memory
  * 50GB available storage
* The following software setup is required
  * docker daemon
  * a non-root account with write access to the docker socket.

## Testground-supplied containers

When testground runs its first plan, several additional containers will be downloaded and started. Here is an overview of everything testground adds to your system

* Docker containers:
  * redis
    * This is part of the testground sync service.
  * grafana
    * Dashboard service
  * Influxdb
    * time series database

