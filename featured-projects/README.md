# Featured projects

## libp2p
libp2p is an open source networking stack and library modularized out of the IPFS Project, and bundled separately for other tools to use.

It features [Testground plans](https://github.com/libp2p/test-plans/) for the overall project, and also [further Testplans](https://github.com/sigp/gossipsub-testground) specifically for the GossipSub pubsub system.


## Testground learning example
This is a pratical example of a "realistic" project, with its own internal business logic, dependencies, etc. The logic and behaviors are intended to be as straightforward as possible, acting as a reference and guide on how to implement behaviors and test them using Testground.

Find the project in the following Github repositories:
 - [Learning project](https://github.com/testground/learning-example)
 - [Testground plans](https://github.com/testground/learning-example-tg)

##### Featured test cases
- Running a build with the `docker:generic` builder, a custom `Dockerfile` and `manifest.toml`
- Running an additional docker container as a dependency for tests (see the `Makefile` for an example)
- Test cases which are intentionally written to fail, due to a context timeout, or a network routing policy that causes a node to be unreachable