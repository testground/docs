# Test plans

## What is a test plan?

A **test plan** is a collection of test cases that exercise, benchmark, or verify a particular component, subsystem, or API of the system under test. 

**Test plans are the unit of deployment that Testground deals with.**

Each test plan is a world of its own. ****In other words, **test plans are** _**opaque**_ **to Testground:** _****_**they behave like black boxes.**

{% hint style="info" %}
Testground does not care what the test plan actually does, the language it's written in, nor the runtime it targets. As long as the source code is accessible, and a builder exists to compile that code into a runnable artefact, such as an executable or a Docker image, you're good to go 🚀 

At the time of writing, Testground offers two builders:

* **`exec:go,`**compiles a Go test plan into a platform executable using the system Go installation.
* **`docker:go,`** compiles a Go test plan into a Docker image.
{% endhint %}

## The test plan &lt;&gt; Testground contract

While test plans are opaque to the eyes of Testground, **test plans and testground promise to satisfy a contract.** That contract is inspired by the [12-factor principles](https://12factor.net/), and facilitates deployment on cloud infrastructure when it's time to scale. The contract is as follows:

1. **Execution:** Test plans expose a single point of entry, i.e. a `main()` function.
2. **Input:** Test plans consume a [formal, standardised runtime environment](../concepts-and-architecture/runtime-environment-runenv.md), in the form of environment variables.
3. **Output:** Test plans record events, results, and optional diagnostics in a predefined JSON schema on stdout and specific files. Any additional output assets they want harvested \(e.g. event trails, traces, generated files, etc.\) are written to a path received in the runtime environment.

{% hint style="success" %}
The Testground community offers SDKs that make it easy for user-written test plans to adhere to the test plan contract, as well as facilitating interactions with the sync service and the emission of metrics.
{% endhint %}

## Test plan manifest

Every test plan must contain a `manifest.toml` file at its root. This is a specification file that declares:

* the name of the test plan, and authorship metadata.
* which builders and runners it can work with.
* the test cases it contains.
* for each test case, the parameters it accepts, including the data type, default value, and a description.

The `manifest.toml` is used by tools such as the **testground CLI,** or the upcoming Jupyter Notebooks integration, to enable a better user experience.  Without this manifest file, it would be impossible to know the contents and behaviour of a test plan without inspecting its source.

For more information on the format of the manifest, see [Writing test plans &gt; Test plan manifest](../writing-test-plans/test-plan-manifest.md).

## Where do test plans live?

Test plans can be hosted and version anywhere—either on the local filesystem, or on public or private Git repositories: Testground does not care.

What's important is that **the source is available to the Testground client during runtime**, under the `$TESTGROUND_HOME/plans` directory, where `$TESTGROUND_HOME` defaults to `$HOME/testground` if not set.

The testground client CLI offers a series of simple commands to manage test plan sources. Refer to the [Managing test plans](../managing-test-plans.md) section for more information.

