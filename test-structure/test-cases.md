# Test cases

## What is a test case?

While the unit of deployment in Testground is the test plan, **test plans nest** _**one or many test cases**_ **inside.**

Think of a test plan as a family of tests that, together, exercises a given component or subsystem. Test cases evaluate concrete use cases that we wish to reproduce consistently, in order to capture variations in the observed behaviour, as the source of the component under test changes over time.

Testground offers first-class support for dealing with test cases inside test plans:

1. **When inspecting a test plan,** the testground CLI allows you to enumerate all test cases within all test plans: `testground plan list --testcases`
2. **When scheduling a test run,** the testground CLI allows you to specify the test case out of a test plan that you want to run, e.g.:

   ```bash
   testground run single --plan libp2p/dht --testcase find-peers
   ```

3. **When developing a test plan,** the testground SDK allows you to select the test case that will run, based on an environment variable.

