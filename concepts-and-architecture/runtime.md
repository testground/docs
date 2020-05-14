---
description: Interacting with the test plan runtime environment
---

# Runtime environment \(runenv\)

Information about the test plan being run is sent to every test instance and available through the `runtime.RunEnv` value - the **runtime environment**.

Let's have a look at the information available to us via the runtime environment:

```go
// RunEnv encapsulates the context for this test run.
type RunEnv struct {
	RunParams
	*logger

  ...
}

// RunParams encapsulates the runtime parameters for this test.
type RunParams struct {
	TestPlan string `json:"plan"`
	TestCase string `json:"case"`
	TestRun  string `json:"run"`

	TestRepo   string `json:"repo,omitempty"`
	TestCommit string `json:"commit,omitempty"`
	TestBranch string `json:"branch,omitempty"`
	TestTag    string `json:"tag,omitempty"`

	TestOutputsPath string `json:"outputs_path,omitempty"`

	TestInstanceCount  int               `json:"instances"`
	TestInstanceRole   string            `json:"role,omitempty"`
	TestInstanceParams map[string]string `json:"params,omitempty"`

	TestGroupID            string `json:"group,omitempty"`
	TestGroupInstanceCount int    `json:"group_instances,omitempty"`

	// true if the test has access to the sidecar.
	TestSidecar bool `json:"test_sidecar,omitempty"`

	// The subnet on which this test is running.
	//
	// The test instance can use this to pick an IP address and/or determine
	// the "data" network interface.
	TestSubnet    *IPNet    `json:"network,omitempty"`
	TestStartTime time.Time `json:"start_time,omitempty"`
}

```

The runtime environment is propagated on any runner to the test instances via environment variables and then deserialized in the `runtime` package of the SDK upon start.

Most of these fields should look familiar -- they represent the configuration we added to our `manifest.toml` file along with configuration specific to the test run and the test instance, such as `TestRun`, or `TestSubnet`.

If you want to find out how many instances of a test plan are running, it's as easy as this:

```go
select runenv.TestInstanceCount {
case 1:
  // The lonliest number
case 2:
  // It takes two to tango
case 3:
  // Do we have a quorum?
default:
  // Now this is a party!
}
```

Additionally a `logger` is initialized as part of the `RunEnv` run environment, so that every time you call `runenv.RecordMessage("my message")` in your test plan, messages include additional metadata such as the current run identifier, group, timestamp, etc.

For more information, check the [Observability, assets and metrics](../writing-test-plans/observability-assets-and-metrics.md).

