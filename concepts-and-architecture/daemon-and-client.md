# Daemon and client

The testground runtime revolves around a traditional daemon/client architecture, communicating over HTTP.

This architecture is flexible enough to run the daemon within a shared cluster, with many users, developers, and integrations \(e.g. GitHub Actions\) hitting the same daemon to schedule build and run jobs.

{% hint style="info" %}
At this moment we are still not running Testground in such shared-cluster deployments; most developers spin up dedicated k8s clusters, and run both the daemon and the client on their own development machines, pointing the local daemon to the remote k8s API.

However, as of Testground v0.5, **we do support running the daemon remotely within a k8s cluster.** This capability is vital to pave the way for shared-cluster deployments, which will enable much more efficient resource utilisation and CI workflows.
{% endhint %}

{% hint style="warning" %}
At the moment, the testground binary contains all three main processes: the daemon, the client, and the sidecar. In the future, we may consider splintering them into separate binaries to save footprint on client-only use cases.
{% endhint %}

## Testground daemon

The testground daemon is responsible for:

1. executing test plan builds.
2. executing test case runs.
3. collecting the outputs from test runs into an archive.
4. performing builder and runner healthchecks.
5. setting up the dependencies for builders and runners to operate.

You start the testground daemon by running:

```text
$ testground daemon
```

Exit with Control-C.

## Testground client

The testground client offers a CLI-based user experience, allowing you to:

1. manage and describe the test plan and SDK sources the client knows about.
2. interact with the daemon by executing builds and runs, collecting outputs, requesting runner termination, and triggering healthchecks.

```text
$ testground help
```

