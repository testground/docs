# Featured projects

## libp2p
libp2p is an open source networking stack and library modularized out of The IPFS Project, and bundled separately for other tools to use.
You can find the libp2p Testground plans in the [Testground Plans Repository](https://github.com/libp2p/test-plans/).

## Testground learning example
This project is intended as a practical example of a "real" project, with its own internal business logic, dependencies, etc. The logic and behaviors are intended to be as straightforward as possible, acting as a reference / guide on how to implement these behaviors and test them using Testground.

You can find the project in the following Github repositories:
 - [Learning project](https://github.com/testground/learning-example)
 - [Testground plans](https://github.com/testground/learning-example-tg)

##### Featured test cases
- Running a build with the `docker:generic` builder, a custom `Dockerfile` and `manifest.toml`
- Running an additional docker container as a dependency for tests (see the `Makefile` for an example)
- Test cases which are intentionally written to fail, due to a context timeout, or a network routing policy that causes a node to be unreachable