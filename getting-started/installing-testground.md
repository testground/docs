# Installing Testground

## Prerequisites

* [Docker](https://www.docker.com/products/docker-desktop)
* [Go 1.14](https://golang.org/) or higher

## Installation

Currently, we don't distribute binaries, so you will have to build from source.

```bash
$ git clone https://github.com/testground/testground.git

$ cd testground

// build testground and the Docker images, used by the local:docker runner.
$ make install

// start the daemon listening on localhost:8042 by default.
$ testground daemon
```

**`$TESTGROUND_HOME`** is an important directory. If not explicitly set, Testground uses `$HOME/testground` as a default.

The layout of **`$TESTGROUND_HOME`** is as follows:

```text
$TESTGROUND_HOME
 |
 |__ plans              >>> [c] contains test plans, can be git checkouts, symlinks to local dirs, or the source itself
 |    |__ suite-a       >>> test plans can be grouped in suites (which in turn can be nested); this enables you to host many test plans in a single repo / directory.
 |    |    |__ plan-1   >>> source of a test plan identified by suite-a/plan-1 (relative to $TESTGROUND_HOME/plans) 
 |    |    |__ plan-2
 |    |__ plan-3        >>> source of a test plan identified by plan-3 (relative to $TESTGROUND_HOME/plans)
 |
 |__ sdks               >>> [c] hosts the test development SDKs that the client knows about, so they can be used with the --link-sdk option.
 |    |__ sdk-go
 |
 |__ data               >>> [d] data directory  
      |__ outputs
      |__ work
 
[c] = used client-side // [d] = used mostly daemon-side.
```

