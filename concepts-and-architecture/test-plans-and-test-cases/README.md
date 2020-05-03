# Test plans, test cases, test runs

## Test plans

A **test plan** contains the logic to exercise a particular component, subsystem, or API of the system under test. 

**Test plans are the unit of deployment that Testground deals with.**

Each test plan is a world of its own. ****In other words, **test plans are** _**opaque**_ **to Testground:** _****_**they behave like black boxes.**

{% hint style="info" %}
Testground does not care what the test plan actually does, the language it's written in, nor the runtime it targets. As long as the source code is accessible, and a builder exists to compile that code into a runnable artefact, such as an executable or a Docker image, you're good to go 🚀 

At the time of writing, Testground offers two builders:

* **`exec:go,`**compiles a Go test plan into a platform executable using the system Go installation.
* **`docker:go,`** compiles a Go test plan into a Docker image.
{% endhint %}

### Test plan contract

In exchange for the above flexibility, **test plans do promise to satisfy a contract**. That contract is inspired by the [12-factor principles](https://12factor.net/), and facilitates deployment on cloud infrastructure when it's time to scale. 

The contract is as follows:

1. **Execution:** Test plans expose a single point of entry, i.e. a `main()` function.
2. **Input:** Test plans consume a [formal, standardised runtime environment](../runtime-environment-runenv.md), in the form of environment variables.
3. **Output:** Test plans record events, results, and optional diagnostics in a predefined JSON schema on stdout and specific files. Any additional output assets they want harvested \(e.g. event trails, traces, generated files, etc.\) are written to a path received in the runtime environment.

{% hint style="success" %}
The Testground community offers SDKs that make it easy for user-written test plans to adhere to the test plan contract, as well as facilitating interactions with the sync service and the emission of metrics.
{% endhint %}

## Test cases

While the unit of deployment in Testground is the test plan—and a test plan exercises a given component or subsystem—in practice there are many use cases, functionalities, or operations that you'll want to test within that component or subsystem.

Test cases evaluate concrete use cases that we wish to reproduce consistently, in order to capture variations in the observed behaviour, as the source of the component under test changes over time.

**That's why test plans can host one or multiple test cases within it.** 

Testground offers first-class support for dealing with test cases inside test plans:

1. **When scheduling a test run,** the testground CLI allows you to specify the test case out of a test plan that you want to run, e.g.:

   ```bash
   testground run single --plan libp2p/dht --testcase find-peers
   ```

2. **When developing a test plan,** the testground SDK allows you to select the test case that will run, based on an environment variable.

## Test run

Every time we execute a test plan, we generate a test run. And each test run is assigned a unique ID. That ID is used to identify the run when collecting outputs, or exploring results or diagnostics.





Language, library choices and testing methodology must be opaque to the test infrastructure and pipeline. Test plans effectively behave like black boxes. A test plan satisfies a contract. It relies on an environment to be injected, executes its custom testing logic, and produces outputs in a predefined manner.

. **Test plans** will be written, owned and maintained by distinct teams and working groups.

\*\*\*\*

Triggered by GitHub automation, manual action, or a schedule \(e.g. nightly\), the test pipeline will:

1. Check out a specific commit of go-ipfs and js-ipfs \(although our immediate priority is go-ipfs\). We call these **test targets**.
2. Build it if necessary.
3. Schedule test plans to be run against it.

We call each run of the test plans a **test run.** Runs are identified by an auto-incrementing ordinal. A given commit xyz could be tested many times, across runs N1, N2, N3, etc. Performing the above steps steps is the responsibility of the **test scheduler** component.

Deliver a **concrete set of inputs** via env variables, following [12-factor](https://12factor.net/config) principles.

1. **Collect and process its output** \(stdout\): metric recordings, test execution results, binary assets to be archived.

## How do I build a test plan?

## **Where do test plans live?**

Test plans can live anywhere.

Test plans are the unit of deployment that the Testground platform deals 

## Example test plans

in the libp2p project, we have built test plans for these components:

* Distributed Hashtable \(DHT\) 
* NAT

A test plan is the 



