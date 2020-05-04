# Understanding Testground performance on Kubernetes

Testground provides a platform for simulating distributed systems at various scale through its `runner` abstraction. This document describes the expected performance when using Testground with the `cluster:k8s` runner on a Kubernetes cloud environment, built using the playbooks from the [github.com/testground/infra](https://github.com/testground/infra) repo.

Within the Testground team, we use the `benchmarks` and the `storm` testplans to establish a performance baseline for the level of scale developers can expect when running testplans on Testground in different cluster configurations.

## Testing environment

* 1 \* `c5.4xlarge` Kubernetes `master` node
* 200 \* `c5.2xlarge` Kubernetes `worker` nodes \(`testplan` group\)
* 3 \* `c5.2xlarge` Kubernetes `worker` nodes \(`infra` group - used for Redis, InfluxDB, Grafana, etc.\)

`c5.2xlarge` - 8 vCPU ; 16GB RAM ; Up to 10Gbps bandwidth

## Resource distribution

When running large scale tests across multiple VMs, it helps to understand how the host / worker node resources are spread between the various testplan instances.

* CPU - all testplan instance have a CPU `request` set. Kubernetes uses kernel throttling to implement CPU limits, and we currently don't set CPU `limits`. If an application goes above the limit, it gets throttled \(aka fewer CPU cycles\). CPU `requests` are used during scheduling of pods.
* Memory - all testplan instances have a limit set for their memory usage. If a testplan tries to use above that limit, Kubernetes will kill it. Memory limits are easy to detect - you only need to check if your pod's last restart status is OOMKilled.
* Disk - all testplan instances are sharing an AWS EFS file system.
* Network - all testplan instances are sharing the underying network as well as the overlay networks \(Weave and Flannel\), so it is possible for a testplan instance to have an adverse effect on all other instances and even on the host. Because of this it is advisable to monitor network performance \(rpc latencies and error rates\) from the perspective of individual test plans.

### Resources available to testplan instances

Alongside testplan instances on every host, we also have the following components:

* `prometheus-node-exporter` - exporter for machine metrics
* `testground-sidecar` - the daemon that monitors containers on a given host and executes network updates, traffic shaping, etc. for testplan instances
* `genie-plugin` - Kubernetes CNI plugin for choosing pod network during deployment time
* `flannel` - the control network daemon
* `kube-proxy` - the Kubernetes network proxy
* `weave` - the data network daemon

All of these consume resources, so testplan instances have ~6 vCPUs and ~12GB RAM available \(assuming `c5.2xlarge` worker node instance type\)

## Testplan scheduling

With a c5.4xlarge `master` node, usually it takes ~5min. for all 10k testplan instances \(pods\) to be scheduled on a given work node and to be in `Running` state. Depending on various factors \(i.e. number of hosts, whether the testplan image is already downloaded, etc.\) we get around 30-50 pods scheduled per second for a given testplan run.

## Sync service

The `sync service` is an integral part of Testground, used to communicate out-of-band information between testplan instances \(such as endpoint addresses, peer ids, etc.\)

At peak operation we've benchmarked that it can perform ~300k operations per second, allowing for 10k testplan instances to publish their own node information and retrieve the peer ids for all other testplan instances in 6min. \(total of 10^8 operations\).

## Networking

We use Weave for our data plane and Flannel for the control network. All testplan instances connect to each other over the Weave network.

Limits for connections per host - based on the kernel parameters we set on individual worker nodes, we can sustain up to 100k open connections per host \(see [github.com/testground/infra/blob/master/k8s/cluster.yaml](https://github.com/testground/infra/blob/master/k8s/cluster.yaml#L113)\)

Depending on the amount of data sent across the network, TCP dial latencies range between 50ms, and can go up to tens of seconds, so we suggest testplan developers monitor this closely.

## Related articles

\[0\] [https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-updated-april-2019-4a9886efe9c4](https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-updated-april-2019-4a9886efe9c4)

\[1\] [https://kubernetes.io/docs/setup/best-practices/cluster-large/](https://kubernetes.io/docs/setup/best-practices/cluster-large/)

\[2\] [https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)

