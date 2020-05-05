# Troubleshooting

Testground is still in early stage of development, so it is possible that:

* Testground crashes
* One of the underlying systems that Testground uses crashes \(`Kubernetes`, `weave` , `redis`, etc.\)
* Testground doesn't properly clean-up after a test run
* etc.

Here are a few commands that could be helpful for you to inspect the state of your Kubernetes cluster and clean up after Testground:

### Delete all pods related to a test plan

Delete all pods that have the `testground.plan=dht` label. This is useful in case you used the `--run-cfg keep_service=true` setting on Testground.

```text
$ kubectl delete pods -l testground.plan=dht --grace-period=0 --force
```

### Restart the sidecar

Restart the `sidecar` daemon which manages networks for all testplans

```text
$ kubectl delete pods -l name=testground-sidecar --grace-period=0 --force
```

### Review running, completed, failed pods

You can check all running pods with

```text
$ kubectl get pods -o wide
```

Another useful combination is `watching` for pods that are not in `Running` state or that are failing their health  checks, with:

```bash
# watch all non-running pods
watch 'kubectl get pods --all-namespaces -o wide | grep -v Running'

# watch all not-ready pods
watch 'kubectl get pods --all-namespaces -o wide | grep "0\/1"'
```

### Get logs from a given pod

```text
$ kubectl logs <pod-id, e.g. tg-dht-c95b5>
```

### Get access to the Redis shell

```text
$ kubectl port-forward svc/testground-infra-redis-master 6379:6379 &
$ redis-cli -h localhost -p 6379
```

