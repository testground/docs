# docker:go builder

The `docker:go` builder uses the user's local docker daemon to construct a docker image. By default, the local:docker builder will leverage a goproxy container to speed up fetching go modules. Additionally, all builds are performed on an isolated docker network.

[github](https://github.com/ipfs/testground/blob/master/pkg/build/golang/docker.go#L40)

| parameter | explanation |
| :--- | :--- |
| go\_version | override the version of golang used to compile the plan |
| module\_path | use an  alternative gomod path in the container |
| exec\_pkg | specify the package name |
| fresh\_gomod | remove and recreate `go.mod` files |
| push\_registry | after build, push docker image to a remote registry |
| registry\_type | must be set if push\_registry is true. Set to `aws` or `dockerhub` |
| go\_proxy\_mode | how to access go proxy. By default, use a local container. |
| go\_proxy\_url | required if `go_proxy_mode` is custom. Use a custom go\_proxy instance. |

