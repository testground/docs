# Paramaters and TestCases

## Adding an additional TestCase

Parameters and TestCases are defined in the come from the manifest toml file. Lets have a look once again at the quickstart toml file and add this to the bottom of the file. 

{% code title="manifests/quickstart.toml" %}
```
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

We would like to test the effect of these two different cases. First, lets prepare our code. Open our quickstart plan program and deal with these parameters!

{% code title="plans/quickstart/main.go" %}
```bash
package main

import (
	"github.com/ipfs/testground/sdk/runtime"
)

func main() {
	runtime.Invoke(run)
}

func run(runenv *runtime.RunEnv) error {
	runenv.Message("I am a %s plan.", runenv.TestCase)
	runenv.Message("I store my files on %d servers.", runenv.IntParam("num"))
	runenv.Message("I %s run tests on my P2P code.", runenv.StringParam("word"))
	if runenv.BooleanParam("feature") {
		runenv.Message("I use IPFS!")
	}
	return nil
}

```
{% endcode %}

Before we run this plan, we can see that the server has two test cases in the listing. We can select which test case we would like to test by runtime flag. With the testground daemon running in another terminal, execute the following to see a list of the new test cases we have just created.

```text
./testground list | grep quickstart
quickstart/smallbrain
quickstart/bigbrain
```

The time has come now to compare these test cases. Lets run it!

```text
./testground run single quickstart/smallbrain \
    --builder exec:go \
    --runner local:exec \
    --instances 1
```

{% hint style="info" %}
Try using different runners. This command executes the plan with he local:exec runner and builder, but it works just as well using your local docker daemon or kubernetes!
{% endhint %}

