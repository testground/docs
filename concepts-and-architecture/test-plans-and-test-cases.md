# Test plans, test cases, test runs

A **test plan** contains the logic to exercise a particular component, subsystem, or API of the system under test. 

**Test plans are the unit of deployment that Testground deals with.**

Each test plan is a world of its own. ****In other words, **test plans are** _**opaque**_ **to Testground:** _****_**they behave like black boxes.**

Testground does not care what the test plan actually does, the language it's written in, nor the runtime it targets. As long as the source code is accessible, and a builder exists to compile that code into a runnable artefact, such as an executable or a Docker image, you're good to go 🚀

However, in exchange for the above flexibility, **test plans do promise to satisfy a contract**. That contract is inspired by the [12-factor principles](https://12factor.net/), and facilitates deployment on cloud infrastructure when it's time to scale. 

The contract is as follows:

1. **Execution:** Test plans expose a single point of entry, i.e. a `main()` function.
2. **Input:** Test plans consume a [formal, standardised runtime environment](runtime-environment-runenv.md), in the form of environment variables.
3. **Output:** Test plans emit events, results, and optional diagnostics in a predefined JSON schema on stdout and specific files. Any additional output assets they want harvested \(e.g. event trails, traces, generated files, etc.\), must be written to a path received in the runtime environment.



t relies on an environment to be injected, executes its custom testing logic, and produces outputs in a predefined manner.

{% hint style="info" %}
At the time of writing, Testground offers two builders:

* **`exec:go,`**which takes a test plan written in Go and compiles it into a platform executable using the system Go SDK.
* **`docker:go,`**which takes a test plan written in Go, and builds a Docker image with the entrypoint set to the main function.
{% endhint %}

## Test cases

Test plans can contain one or multiple **test cases** within it. Test cases evaluate concrete use cases that we wish to reproduce consistently over time, in order to capture variations in the observed behaviour, as the source under test evolves.

Every instance of a test plan execution is called a **test run,** and it is assigned a unique ID in the system.





Language, library choices and testing methodology must be opaque to the test infrastructure and pipeline. Test plans effectively behave like black boxes. A test plan satisfies a contract. It relies on an environment to be injected, executes its custom testing logic, and produces outputs in a predefined manner.

. **Test plans** will be written, owned and maintained by distinct teams and working groups.

\*\*\*\*

Triggered by GitHub automation, manual action, or a schedule \(e.g. nightly\), the test pipeline will:

1. Check out a specific commit of go-ipfs and js-ipfs \(although our immediate priority is go-ipfs\). We call these **test targets**.
2. Build it if necessary.
3. Schedule test plans to be run against it.

We call each run of the test plans a **test run.** Runs are identified by an auto-incrementing ordinal. A given commit xyz could be tested many times, across runs N1, N2, N3, etc. Performing the above steps steps is the responsibility of the **test scheduler** component.

For each test plan, the test runtime _promises_ to:

1. Deliver a **concrete set of inputs** via env variables, following [12-factor](https://12factor.net/config) principles.
2. **Collect and process its output** \(stdout\): metric recordings, test execution results, binary assets to be archived.

## How do I build a test plan?

## **Where do test plans live?**

Test plans can live anywhere.

Test plans are the unit of deployment that the Testground platform deals 

## Example test plans

in the libp2p project, we have built test plans for these components:

* Distributed Hashtable \(DHT\) 
* NAT

A test plan is the 



