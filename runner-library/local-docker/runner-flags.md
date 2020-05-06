# Runner flags

The `local:docker`runner accepts flags on the command-line which can modify its behavior. The chart below shows available options:

In order to pass a command-line flag, add the option `--run-cfg` followed by any of the following options:

[github](https://github.com/ipfs/testground/blob/master/pkg/runner/local_docker.go#L49)

| parameter | explanation |
| :--- | :--- |
| keep\_containers | specify whether containers should be removed after execution |
| log\_level | specify the logging verbosity |
| no\_start | if set, containers will be created but not executed |
| background | if set, the output of containers will not be displayed |
| ulimits | override ulimits applied to docker containers |

## 

