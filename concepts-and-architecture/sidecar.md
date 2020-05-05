# Sidecar

**The sidecar runs in pure Docker and Kubernetes environments** \(i.e. the `local:docker` and `cluster:k8s`runners\). It does not run on `local:exec`. 

**It is one of the three processes in the Testground executable;** the other two are the daemon and the client. It can be started with:

```bash
$ testground sidecar --runner local:docker # or cluster:k8s
```

On Kubernetes, each node runs a copy of the sidecar. We schedule it via a [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/). 

Usually, you will never need to start the sidecar manually. The runner healthchecks \(which run prior to a test run\) will ensure that the appropriate sidecars are started.

### Responsibilities

Sidecars are responsible for three things:

1. Initializing the network interfaces of test instances. See [Networking](networking.md) for more info.
   * Sidecars watch the local Docker daemon for containers being scheduled.
   * For each new container, they adjust the routing of the `control` network, and they initialize the `data` network, incrementing the `network-initialized` state on the sync service every time that a new container is instrumented under that `run_id`.
2. Applying networking configurations requested by the container, by targeting the appropriate data interface adaptor through [Netlink](http://man7.org/linux/man-pages/man7/netlink.7.html). See [Networking](networking.md) for more info.
3. Periodically garbage-collecting inactive entries in the [sync service](sync-service.md) \(Redis\), pertaining to finished runs.

