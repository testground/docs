# docker:go

## Background

Provided that you are writing plans in go and using the [sync service sdk](https://github.com/testground/sdk-go),  `docker:go` is the builder you want to use. This builder produces plan artifacts which are usable on the `local:docker`  or `cluster:k8s` runner.

## Features

* Produces docker containers which can run on docker-powered runners.
* Linkage with a custom sync service SDK via `--link-sdk` flag
* Simple command-line dependency mapping uisng the `--dep` flag
* Busybox base for easy container troubleshooting.

## Troubleshooting

Most build failures are caused by a problem with the plan code rather than the build process. Here are a few tips to help figure out what's going on when there are build failures.

1. Try to build the plan yourself. These are just executable files after all! Frequently, build problems can be revealed by simply trying to run the plan as it is. If you can't build it like this, then the builder won't be able to either.

   ```text
   cd <plan_directory>
   export GOMOD111MODULE=on go build .
   ```

2. Make sure you have correctly initialized `go.mod`. The builder may throw errors when doing module replacements not used correctly.
3. Remove any custom module replace lines you have in `go.mod`. The builder will do this for you with appropriate flags.

## Learn More

Interested in how this works? All the testground builders can be seen [here](https://github.com/testground/testground/tree/master/pkg/build)



