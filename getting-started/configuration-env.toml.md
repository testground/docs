# Configuration \(.env.toml\)

`.env.toml`is a configuration file read by the Testground daemon and Testground client on startup.

Testground tries to load this file from `$TESTGROUND_HOME/.env.toml`

### Changing default bind addresses

You can change the default bind addresses by configuring `daemon.listen` and `client.endpoint`

```text
[daemon]
listen = ":8080"

[client]
endpoint = "localhost:8080"
```

### AWS integration

When using a remote runner such as `cluster:k8s`, you should configure the default region:

```text
["aws"]
region = "aws region, such as eu-central-1"
```

### DockerHub integration

If you want to push Docker images from the `docker:go` builder to a remote registry, you can configure it.

```text
["dockerhub"]
repo = "repo to be used for testground"
username = "username"
access_token = "docker hub access token"
```



