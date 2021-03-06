# Understanding Testground performance on Kubernetes

## Resource distribution

When running large scale tests across multiple VMs, it helps to understand how the host / worker node resources are spread between the various test plan instances.

* CPU - all test plan instance have a CPU `request` set. Kubernetes uses kernel throttling to implement CPU limits, and we currently don't set CPU `limits`. If an application goes above the limit, it gets throttled \(aka fewer CPU cycles\). CPU `requests` are used during scheduling of pods.
* Memory - all test plan instances have a limit set for their memory usage. If a test plan tries to use above that limit, Kubernetes will kill it. Memory limits are easy to detect - you only need to check if your pod's last restart status is `OOMKilled`.
* Disk - all test plan instances are sharing an AWS EFS file system.
* Network - all test plan instances are sharing the underlying network as well as the overlay networks \(Weave and Flannel\), so it is possible for a test plan instance to have an adverse effect on all other instances and even on the host. Because of this it is advisable to monitor network performance \(RPC latencies and error rates\) from the perspective of individual test plans.

### Resources available to test plan instances

Alongside test plan instances on every host, we also have the following components:

* `prometheus-node-exporter` - exporter for machine metrics
* `testground-sidecar` - the daemon that monitors containers on a given host and executes network updates, traffic shaping, etc. for test plan instances
* `genie-plugin` - Kubernetes CNI plugin for choosing pod network during deployment time
* `flannel` - the control network daemon
* `kube-proxy` - the Kubernetes network proxy
* `weave` - the data network daemon

All of these consume resources, so test plan instances have ~6 vCPUs and ~12GB RAM available \(assuming `c5.2xlarge` worker node instance type\)

## Test plan scheduling

With a c5.4xlarge `master` node, usually it takes ~5min. for all 10k test plan instances \(pods\) to be scheduled on a given work node and to be in `Running` state. Depending on various factors \(i.e. number of hosts, whether the test plan image is already downloaded, etc.\) we get around 30-50 pods scheduled per second for a given test plan run.

## Sync service

The `sync service` is an integral part of Testground, used to communicate out-of-band information between test plan instances \(such as endpoint addresses, peer ids, etc.\)

At peak operation we've benchmarked that it can perform ~300k operations per second, allowing for 10k test plan instances to publish their own node information and retrieve the peer ids for all other test plan instances in 6min. \(total of 10^8 operations\).

## Networking

We use Weave for our data plane and Flannel for the control network. All test plan instances connect to each other over the Weave network.

Limits for connections per host - based on the kernel parameters we set on individual worker nodes, we can sustain up to 100k open connections per host \(see [github.com/testground/infra/blob/master/k8s/cluster.yaml](https://github.com/testground/infra/blob/master/k8s/cluster.yaml#L113)\)

Depending on the amount of data sent across the network, TCP dial latencies range between 50ms, and can go up to tens of seconds if you have thousands of calls, so we suggest test plan developers monitor this closely.

## Related articles

\[0\] [https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-updated-april-2019-4a9886efe9c4](https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-updated-april-2019-4a9886efe9c4)

\[1\] [https://kubernetes.io/docs/setup/best-practices/cluster-large/](https://kubernetes.io/docs/setup/best-practices/cluster-large/)

\[2\] [https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)

