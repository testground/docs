---
description: Interacting with the environment.
---

# Runtime

## Getting information about the test run

Following the [https://12factor.net/config](https://12factor.net/config) methodology, variable information and configuration is stored in the environment. Information about the plan being run as well as any parameters passed just one `os.Getenv()` away. 

Lets have a look at the information available to us as environment variables. Edit the quickstart plan so it looks like the following:

```go
package main

import (
	"github.com/ipfs/testground/sdk/runtime"
	"os"
)

func main() {
	runtime.Invoke(run)
}

func run(runenv *runtime.RunEnv) error {
	for _, env := range os.Environ() {
		runenv.RecordMessage(env)
	}
	return nil
}

```

When this plan is executed \(on any runner\) you should notice that several environment variables are passed. Most of these should look familiar -- these variables represent the configuration we added to our toml manifest along with configuration specific to the test run and and test instance.

```text
TEST_GROUP_ID=single
TEST_INSTANCE_COUNT=1
TEST_RUN=442fe7759bf8
TEST_BRANCH=
TEST_CASE_SEQ=0
TEST_SIDECAR=false
TEST_PLAN=quickstart
TEST_START_TIME=2020-04-09T17:45:31-07:00
TEST_GROUP_INSTANCE_COUNT=1
TEST_OUTPUTS_PATH=/some/path/outputs/quickstart/442fe7759bf8/single/0
TEST_INSTANCE_ROLE=
TEST_REPO=
TEST_SUBNET=127.1.0.0/16
TEST_TAG=
TEST_CASE=testcase1
TEST_INSTANCE_PARAMS=who="world"
```

As a convenience, these environment variables are added to the RunParams struct, For example, if you want to find out how many instances of a plan are running, it's as easy as this:

{% code title="" %}
```go
select runenv.TestInstanceCount {
case 1:
  // The lonliest number
case 2:
  // It takes two to tango
default:
  // Lets party!
}
```
{% endcode %}



