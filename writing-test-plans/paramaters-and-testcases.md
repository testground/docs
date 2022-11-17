# Parameters and test cases

## Adding an additional test case

Parameters and test cases are defined in the `manifest.toml` file of every test plan. Let's have a look once again at the `quickstart/manifest.toml` file and add two test cases to the bottom of the file:

{% code title="$TESTGROUND\_HOME/plans/quickstart/manifest.toml" %}
```text
...


[[testcases]]
name = "smallbrain"
instances = { min = 1, max = 200, default = 1 }

  [testcases.params]
  word = { type = "string", default = "never" }
  num = { type = "int", default = 2 }
  feature = { type = "bool", default = false }

[[testcases]]
name = "bigbrain"

  [testcases.params]
  word = { type = "string", default = "always" }
  num = { type = "int", default = 10000000 }
  feature = { type = "bool", default = false }
```
{% endcode %}

?> Feel free to add your own `galaxybrain` test case as well!

You can confirm that the [testground client](../concepts-and-architecture/daemon-and-client.md#testground-client) is able to parse the manifest, and enumerate the cases it declares:

```bash
$ testground plan list --testcases | grep quickstart
quickstart    smallbrain
quickstart    bigbrain
```

We would like to take these two cases out for a spin. First, let's prepare our code. Open our `quickstart` plan program and deal with these parameters!

!> This snippet is routing both test cases to the same function. In practice, you will want to run different logic for each test case!

As you can see, a test plan a simple executable that conforms to the [simple Testground contract](../concepts-and-architecture/test-structure.md#the-test-plan-less-than-greater-than-testground-contract), which the SDK facilitates. This makes it super easy to debug and develop. There's beauty in simplicity!

{% code title="$TESTGROUND\_HOME/plans/quickstart/main.go" %}
```go
package main

import (
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

func main() {
	run.InvokeMap(map[string]interface{}{
		"bigbrain":   test,
		"smallbrain": test,
	})
}

func test(runenv *runtime.RunEnv) error {
	var (
		num     = runenv.IntParam("num")
		word    = runenv.StringParam("word")
		feature = runenv.BooleanParam("feature")
	)

	runenv.RecordMessage("I am a %s test case.", runenv.TestCase)
	runenv.RecordMessage("I store my files on %d servers.", num)
	runenv.RecordMessage("I %s run tests on my P2P code.", word)

	if feature {
		runenv.RecordMessage("I use IPFS!")
	}

	return nil
}
```
{% endcode %}

The time has come now to run these test cases. Let's run it!

```bash
$ testground run single --plan quickstart \
                        --testcase smallbrain \
                        --builder exec:go \
                        --run local:exec \
                        --instances 1 \
                        --wait
```

?> Try using different runners. This command executes the plan with the `local:exec` runner and `exec:go`builder, but it works just as well with the `local:docker` runner or the Kubernetes `cluster:k8s`runner \(for which you will need to use the `docker:go` builder!
