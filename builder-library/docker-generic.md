# docker:generic

## Background

Sometimes the default builder behaviors just aren't flexible enough to do what you want. For those situations, we have the generic builder. Linking and SDK replacement features in the other builders are missing here. No, to use this builder, you must provide your own Dockerfile and specify all the appropriate build activities yourself. 

## Features

* Produces docker containers which can run on docker-powered runners.
* Supports any programming language
* Add assets, files, or whole applications
* highly customizable
* Docker arguments in toml file or command-line flags.

## Non-features

* No automatic SDK linkage
* No language-specific features i.e. module dependency management.

## Usage

1. Place a Dockerfile inside the plan directory.
2. While writing your Dockerfile, you must recognize that the root of the docker build context is one directory higher than the plan. This is the way your plan is packaged when it is sent to the testground daemon. The directory structure actually includes the plan and the SDK
3. With this directory structure in mind, the builder will construct the container using the equivilent of this command: `docker build -f buildctx/plan/Dockerfile buildctx`. Keep that in mind as you consider which paths to use.
4. An example which uses docker arguments and adds additional files can be seen [here](https://github.com/testground/testground/tree/master/plans/example).

```text
buildctx/
├── plan                          <- Your plan is always in a /plan directory
│   ├── Dockerfile                <- Your provided Dockerfile
│   └── yourprogram.sourcecode    <- program files, {c, go, rs, py, java, etc}
└── sdk                           <- sdk code which comes with --link-sdk flag
```

## Running the example

This example runs a  test case which depends on an additional file being added to the image -- something that would not be easy to do using the docker:go builder. Have a look at the Dockerfile and the manifest.toml files to see how this is constructed.

```text
testground plan import --from $GOPATH/src/github.com/testground/testground/plans/example
testground run single --builder docker:generic --runner local:docker --plan example --testcase artifact
```

## Are you using docker:generic?

If you are using this, we think that's awesome! Maybe there are cool features which we should add to our other builders or interesting additions we haven't considered. We want to hear about it! Open an issue with us on github and let us know what your're up to!

* Are you using an unsupported language?
* Are you testing with frameworks that aren't easily achievable using the other language-specific builders? \(I'm looking at you, headless web browsers!\)
* Are there bugs in the other builders you are circumventing? Definitely let us know if this is the case.

## Learn More

Interested in how this works? All the testground builders can be seen [here](https://github.com/testground/testground/tree/master/pkg/build)

