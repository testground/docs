# Daemon and client

The Testground runtime revolves around a traditional daemon/client architecture, communicating over HTTP.

This architecture is flexible enough to run the daemon within a shared cluster, with many users, developers, and integrations \(e.g. GitHub Actions\) hitting the same daemon to schedule build and run jobs.

{% hint style="info" %}
At the moment, we do not run Testground in shared-cluster deployments; most developers spin up personal Kubernetes clusters, and run both the daemon and the client on their own development machines, pointing the local daemon to the remote k8s API.

However, as of Testground v0.5, **we do support running the daemon within a Kubernetes cluster as a pod.** This capability is vital to pave the way for shared-cluster deployments, which will enable much more efficient resource utilization and CI workflows.
{% endhint %}

## Testground Daemon

The Testground daemon is responsible for:

1. Scheduling and executing tasks \(plan builds and test runs\).
2. Collecting the outputs from test runs into an archive.
3. Performing builder and runner healthchecks.
4. Setting up the dependencies for builders and runners to operate.

You start the Testground daemon by running:

```text
$ testground daemon
```

Exit with `Control-C`.

## Testground Client

The Testground client offers a CLI-based user experience, allowing you to:

1. Manage and describe the test plan and SDK sources the client knows about.
2. Interact with the daemon by executing builds and runs, collecting outputs, requesting runner termination, and triggering healthchecks.

```text
$ testground help
```

## Asynchronous Tasks

The Testground daemon executes tasks, i.e., plan builds and test runs, asynchronously. By default, when you execute a build or run command, you'll be returned a task ID. This ID can be used to check the status of the task:

```text
$ testground status -t <id>
```

Or see the task logs - you can use `-f` to follow the logs live:

```text
$ testground logs -t <id> [-f]
```

To visualize all the running tasks:

```text
$ testground tasks
```

You can also wait for the task completion when building or running by appending the flag `--wait`. This way, you will not only wait for the completion of the task, but also receive all the logs live:

```text
$ testground build single --plan=example --builder=exec:go --wait
```

