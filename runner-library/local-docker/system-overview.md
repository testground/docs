# System overview

## Background

Using the `local:docker` runner of Testground enables you to test distributed/p2p on your local machine precisely the same build artifact as the `cluster:k8s` runner. The use of docker images as a build target ensures plans are built the same way with the same build environment even when built on a different machine or a different OS.

The overall performance of the `local:docker` runner depends highly on the machine running the tests, though we are frequently able to run tests with 100+ instances without issue using commonly-available laptop or desktop computers.

## Dependencies:

Aside from the sidecar, which must be built separately, any auxiliary containers are downloaded if needed. This is what you need to get started:

* A laptop or desktop with reasonable hardware specs. If your computer is newer than 5 or 6 years old, it will probably be just fine. To be able run simulations of a reasonable size, we recommend at least the following:
  * 8GB memory
  * 50GB available storage
* The following software setup is required
  * docker daemon
  * a non-root account with write access to the docker socket.
  * Sidecar container installation. See installation instructions in `getting started`

## Testground-supplied containers

When testground runs its first plan, several additional containers will be downloaded and started. Here is an overview of everything testground adds to your system

* Docker networks:
  * testground-build
    * This network is used just during the build phase.
    * Plan containers are not attached to this network.
  * testground-control
    * This network is used by plan containers to communicate with the sync service and the host.
    * Though this is a public-facing network, plan containers cannot access the internet through this network because routes are removed.
  * testground-data
    * Plan containers communicate with each other over this network.
    * This is the network used for testing.
    * The sidecar modifies the link quality and performance to to simulate real-world conditions.
* Docker containers:
  * goproxy
    * This container is attached to the build network.
    * Speeds up golang builds
    * Optional.
  * redis
    * This is part of the testground sync service.
  * grafana
    * Dashboard service
  * Influxdb
    * time series database

