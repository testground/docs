# System overview

## Background

Using the `cluster:k8s` runner of Testground enables you to test distributed/p2p systems at scale.

The `cluster:k8s` Testground runner is capable of launching test workloads comprising 10k+ instances, and we aim to reach 100k at some point.

The [IPFS](https://ipfs.io/) and [libp2p](https://libp2p.io/) projects have used these scripts and playbooks to deploy large-scale test infrastructure. By crafting test scenarios that exercise components at such scale, we have been able to run simulations, carry out attacks, perform benchmarks, and execute all kinds of tests to validate correctness and performance.

## Dependencies

### Kubernetes Operations \(kops\)

Kubernetes Operations \(`kops`\) is a tool which helps to create, destroy, upgrade and maintain production-grade Kubernetes clusters from the command line. We use it to create a Kubernetes cluster on AWS.

### CoreOS Flannel

We use CoreOS Flannel for networking on Kubernetes - for the default Kubernetes network, which in Testground terms is called the `control` network.

`kops` uses 100.96.0.0/11 for pod CIDR range, so this is what we use for the `control` network.

### Weave Net

We use Weave Net for the `data` plane on Testground - a secondary overlay network that we attach containers to on-demand.

We configure Weave to use 16.0.0.0/4 as CIDR \(we want to test `libp2p` nodes with IPs in public ranges\), so this is the CIDR for the Testground `data` network. The `sidecar` component is responsible for setting up the `data` network for every testplan instance.

### CNI-Genie

In order to have two different networks attached to pods in Kubernetes, we run the [CNI-Genie CNI](https://github.com/cni-genie/CNI-Genie).

