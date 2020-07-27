# Troubleshooting

## Verify that Docker is working correctly

The following command, if it runs successfully, will verify that docker is running correctly.

```text
$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
0e03bdcc26d7: Pull complete 
Digest: sha256:8e3114318a995a1ee497790535e7b88365222a21771ae7e53687ad76563e8e76
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

## Docker error messages

| message | cause | fix |
| :--- | :--- | :--- |
| `engine build error: docker build failed: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?` | Docker daemon is not running. | Restart the Docker daemon |
| `read tcp xxx.xxx.xxx.xxx:443 i/o timeout` | Cannot connect to remote repository | This is a frequent and probably transient error. Try again. |

## Troubleshooting

* Test plans fail to finish before context expires

  This may be caused by high load on the host machine. Try to reduce the load using the following:

  * Reduce the number of instances
  * terminate old test plan containers \(if any\) with `testground terminate`
  * close applications which may cause enough load to interfere with the test

* Test plans fail while waiting on a barrier

  This may indicate a problem with the sync service or the sidecar. Investigate and try again.

  * `docker logs testground-redis`
  * `docker kill testground-redis`

* Testground containers are killed before plans finished.

  It's possible that the the system is killed after running out of memory or that another system problem is occurring.

  * `docker events`
  * `docker logs <suspicious_container>`
  * `docker stats`

## Healthchecks

Testground comes equipped with `healthchecks` that have _some_ self-fixing features. These healthchecks are able to fix any container that is added by Testground, but it cannot fix anything related to your user account or host system.

To view current `healthcheck` status, and issue automatic remediation, use the following commands. Under normal circumstances, this is unnecessary since they will be started automatically if a problem is detected.

```bash
# view the healthcheck status
$ testground healthcheck --runner local:docker

finished checking runner local:docker
Checks:
- local-outputs-dir: ok; directory exists.
- control-network: failed; network does not exist.
- local-grafana: failed; container not found.
- local-redis: failed; container not found.
- local-influxdb: failed; container not found.
- sidecar-container: failed; container not found.
Fixes:
- control-network: omitted; 
- local-grafana: omitted; 
- local-redis: omitted; 
- local-influxdb: omitted; 
- sidecar-container: omitted; 


# fix the problems automatically
$ testground healthcheck --runner local:docker --fix

... a few seconds later

finished checking runner local:docker
Checks:
- local-outputs-dir: ok; directory exists.
- control-network: failed; network does not exist.
- local-grafana: failed; container not found.
- local-redis: failed; container not found.
- local-influxdb: failed; container not found.
- sidecar-container: failed; container not found.
Fixes:
- local-outputs-dir: unnecessary; 
- control-network: ok; network created.
- local-grafana: ok; container started
- local-redis: ok; container started
- local-influxdb: ok; container started
- sidecar-container: ok; container started
```

