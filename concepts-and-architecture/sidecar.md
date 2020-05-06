# Sidecar

This `sidecar` is an individual Testground process responsible for network management and traffic shaping for test plan instances. It runs in privileged-mode on host machines, and listens for requests from test plan instances for network configuration through the `sync service`.

The `sidecar` is one of the three processes in the Testground executable. The other two are the Testground daemon and the Testground client. It can be started with:

```bash
$ testground sidecar --runner local:docker # or cluster:k8s
```

**The sidecar runs in Docker and Kubernetes environments** \(i.e. when using the `local:docker` or the `cluster:k8s`runners\). For now it is not supported and it does not run when using the `local:exec` runner.

On Kubernetes, each worker node runs the sidecar. We schedule it via a [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/).

Usually, you will never need to start the sidecar manually. Each runner's health checks \(which run prior to a test run\) will ensure that the sidecar is running.

## Responsibilities

Sidecars are responsible for three things:

1. Initializing the network interfaces of test instances. See [Networking](networking.md) for more info.
   * Sidecars watch the local Docker daemon for containers being started and stopped.
   * For each new container, they adjust the routing of the `control` network, and they initialize the `data` network, incrementing the `network-initialized` state on the `sync service` every time a new test plan container is instrumented under a given `run_id`.
2. Applying networking configurations requested by test plan instances, by targeting the appropriate data interface through [Netlink](http://man7.org/linux/man-pages/man7/netlink.7.html). See [Networking](networking.md) for more info.
3. Periodically garbage-collect inactive entries in the [sync service](sync-service.md) \(backed by a Redis database\), pertaining to finished runs.

