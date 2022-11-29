# Observability, assets and metrics

_Gathering data about your test run_

In the [Getting Started](../getting-started.md) section, we showed you how to `collect` test run `outputs`, i.e. the events outputted by individual test run instances during the test run, for example:

```go
	// emit a metric
	runenv.R().RecordPoint("time-to-find", elapsed.Seconds())

	c := "some GUID"
	// emit a message
	runenv.RecordMessage("provided content ID: %s", c)
```

Testground SDK supports a variety of metrics and outputs to support test plan developers in gathering data and understanding the results of their test plans and test runs.

## Availability

Events outputted by test instances via the available APIs from the `RunEnv` runtime environment, are generally available as:

* `output` files after the test run concludes
* in the metrics database, currently InfluxDB

## Lifecycle events

Lifecycle events facilitate real-time progress monitoring of test runs, either by a human, or by the upcoming watchtower/supervisor service.

They are inserted immediately into the metrics database via direct call to the InfluxDB API.

### API

###### sdk-go/runtime/runtime\_events.go
```text
runenv.RecordStart()
runenv.RecordFailure(reason)
runenv.RecordCrash(err)
runenv.RecordSuccess()
```

## Diagnostics

Diagnostics are inserted immediately into the metrics store via direct call to the InfluxDB API. They are also recorded in file `diagnostics.out`.

### API

###### sdk-go/runtime/metrics\_api.go
```text
runenv.D().RecordPoint(name, value, unit)
runenv.D().Counter()
runenv.D().EWMA()
runenv.D().Gauge()
runenv.D().GaugeF()
runenv.D().Histogram()
runenv.D().Meter()
runenv.D().Timer()
runenv.D().SetFrequency()
runenv.D().Disable()
```

## Results

Recording observations about the subsystems and components under test. Conceptually speaking, results are a part of the test output.

Results are the end goal of running a test plan. Results feed comparative series over runs of a test plan, along time, across dependency sets.

They are batch-inserted into InfluxDB when the test run concludes.

### API

###### sdk-go/runtime/metrics\_api.go
```text
runenv.R().RecordPoint(name, value, unit)
runenv.R().Counter()
runenv.R().EWMA()
runenv.R().Gauge()
runenv.R().GaugeF()
runenv.R().Histogram()
runenv.R().Meter()
runenv.R().Timer()
runenv.R().SetFrequency()
```

## Assets

Output assets will be saved when the test terminates. You can also manually create output assets/directories under `runenv.TestOutputsPath`.

### API

###### sdk-go/runtime/runenv\_assets.go
```text
runenv.CreateRandomDirectory(directoryPath string, depth uint)
runenv.CreateRandomFile(directoryPath string, size int64)
runenv.CreateRawAsset(name string)
runenv.CreateStructuredAsset(name string, config zap.Config)
```
