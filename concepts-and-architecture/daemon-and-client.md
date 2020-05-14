# Daemon and client

The Testground runtime revolves around a traditional daemon/client architecture, communicating over HTTP.

This architecture is flexible enough to run the daemon within a shared cluster, with many users, developers, and integrations \(e.g. GitHub Actions\) hitting the same daemon to schedule build and run jobs.

{% hint style="info" %}
At the moment, we do not run Testground in shared-cluster deployments; most developers spin up personal Kubernetes clusters, and run both the daemon and the client on their own development machines, pointing the local daemon to the remote k8s API.

However, as of Testground v0.5, **we do support running the daemon within a Kubernetes cluster as a pod.** This capability is vital to pave the way for shared-cluster deployments, which will enable much more efficient resource utilization and CI workflows.
{% endhint %}

## Testground daemon

The Testground daemon is responsible for:

1. Executing test plan builds.
2. Executing test case runs.
3. Collecting the outputs from test runs into an archive.
4. Performing builder and runner healthchecks.
5. Setting up the dependencies for builders and runners to operate.

You start the Testground daemon by running:

```text
$ testground daemon
```

Exit with `Control-C`.

## Testground client

The Testground client offers a CLI-based user experience, allowing you to:

1. Manage and describe the test plan and SDK sources the client knows about.
2. Interact with the daemon by executing builds and runs, collecting outputs, requesting runner termination, and triggering healthchecks.

```text
$ testground help
```

