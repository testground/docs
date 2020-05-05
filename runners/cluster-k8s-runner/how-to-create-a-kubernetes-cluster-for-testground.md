# How to create a Kubernetes cluster for Testground

## Requirements

First and foremost, you need an AWS account with API access.

Next, download and install all required software:

1. [kops](https://github.com/kubernetes/kops/releases) &gt;= 1.17.0
2. [terraform](https://terraform.io/) &gt;= 0.12.21
3. [AWS CLI](https://aws.amazon.com/cli)
4. [helm](https://github.com/helm/helm) &gt;= 3.0

## Setup AWS cloud credentials

1. [Generate your AWS IAM credentials](https://console.aws.amazon.com/iam/home#/security_credentials).
2. [Configure the aws-cli tool with your credentials](https://docs.aws.amazon.com/cli/).

## Generate a Testground SSH key for `kops`

It is used for the Kubernetes master and worker nodes

```bash
# generate ~/.ssh/testground_rsa    
#          ~/.ssh/testground_rsa.pub

$ ssh-keygen -t rsa -b 4096 -C "your_email@example.com" \
                            -f ~/.ssh/testground_rsa -q -P ""
```

## Create a bucket for `kops` state

This is similar to Terraform state bucket.

```text
$ aws s3api create-bucket \
      --bucket <bucket_name> \
      --region <region> --create-bucket-configuration LocationConstraint=<region>
```

Where:

* `<bucket_name>` is an AWS account-wide unique bucket name to store this cluster's `kops` state, e.g. `kops-backend-bucket-<your_username>`.
* `<region>` is an AWS region like `eu-central-1` or `us-west-2`.

## Configure cluster variables

* a cluster name \(for example `name.k8s.local`\)
* set AWS region
* set AWS availability zone A \(not region; for example `us-west-2a` \[availability zone\]\) - used for master node and worker nodes
* set AWS availability zone B \(not region; for example `us-west-2b` \[availability zone\]\) - used for more worker nodes
* set `kops` state store bucket \(the bucket we created in the section above\)
* set number of worker nodes
* set location of your cluster SSH public key \(for example `~/.ssh/testground_rsa.pub` generated above\)

You might want to add them to your `rc` file \(`.zshrc`, `.bashrc`, etc.\), or to an `.env.sh` file that you source.

```bash
export NAME=<desired kubernetes cluster name (e.g. mycluster.k8s.local)>
export KOPS_STATE_STORE=s3://<kops state s3 bucket>
export AWS_REGION=<aws region, for example eu-central-1>
export ZONE_A=<aws availability zone, for example eu-central-1a>
export ZONE_B=<aws availability zone, for example eu-central-1b>
export WORKER_NODES=4
export PUBKEY=$HOME/.ssh/testground_rsa.pub
export TEAM=<optional - your team name ; tag is used for cost allocation purposes>
export PROJECT=<optional - your project name ; tag is used for cost allocation purposes>
```

## Setup required Helm chart repositories

```text
$ helm repo add stable https://kubernetes-charts.storage.googleapis.com/
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm repo add influxdata https://helm.influxdata.com/
$ helm repo update
```

## Configure the Testground client

Create a `.env.toml` file in your `$TESTGROUND_HOME` and add your AWS region to the `[aws]` section.

## Create cloud resources for the Kubernetes cluster

This will take about 10-15 minutes to complete.

Once you run this command, take some time to walk the dog, clean up around the office, or go get yourself some coffee! When you return, your shiny new Kubernetes cluster will be ready to run Testground plans.

```text
$ ./k8s/install.sh ./k8s/cluster.yaml
```

## Destroy the cluster when you're done working on it

Do not forget to delete the cluster once you are done running test plans.

```text
$ ./k8s/delete.sh
```

## Resizing the cluster

Edit the cluster state and change number of nodes.

```text
$ kops edit ig nodes
```

Apply the new configuration

```text
$ kops update cluster $NAME --yes
```

Wait for nodes to come up and for `DaemonSets` to be `Running` on all new nodes

```text
$ watch 'kubectl get pods'
```

## Observability

Access to Grafana \(admin credentials are `username: admin` ; `password: admin`\):

```text
$ kubectl port-forward service/testground-infra-grafana 3000:80
```

## Cleanup after Testground and other useful commands

Testground is still in early stage of development.

It is possible that it crashes, or doesn't properly clean-up after a testplan run. Here are a few commands that could be helpful for you to inspect the state of your Kubernetes cluster and clean up after Testground.

* Delete all pods that have the `testground.plan=dht` label \(in case you used the `--run-cfg keep_service=true` setting on Testground.

```text
$ kubectl delete pods -l testground.plan=dht --grace-period=0 --force
```

* Restart the `sidecar` daemon which manages networks for all testplans

```text
$ kubectl delete pods -l name=testground-sidecar --grace-period=0 --force
```

* Review all running pods

```text
$ kubectl get pods -o wide
```

* Get logs from a given pod

```text
$ kubectl logs <pod-id, e.g. tg-dht-c95b5>
```

* Check on the monitoring infrastructure \(it runs in the monitoring namespace\)

```text
$ kubectl get pods --namespace monitoring
```

* Get access to the Redis shell

```text
$ kubectl port-forward svc/testground-infra-redis-master 6379:6379 &
$ redis-cli -h localhost -p 6379
```

## Use a Kubernetes context for another cluster

`kops` lets you download the entire Kubernetes context config.

If you want to let other people on your team connect to your Kubernetes cluster, you need to give them the information.

```text
$ kops export kubecfg --state $KOPS_STATE_STORE --name=$NAME
```

