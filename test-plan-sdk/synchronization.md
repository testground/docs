# Keeping instances in sync

Sometimes individual instances of a test plan need to coordinate what they are doing. For this, we use Barriers.

The general concept is this -- as a plan reaches a phase of execution for which synchronization is required, it will signal the other instances that it has reached that phase. Then, wait until a certain number of other instances reach the same state before continuing.

{% hint style="info" %}
Internally, this is a simple counter backed by a redis database. To wait for a signal barrier means to wait until the counter reaches the specified number.
{% endhint %}

Lets have a look at a plan that waits for others! The plan displayed here is out of the [example plans](https://github.com/ipfs/testground/tree/master/plans/example). Notice the use of `writer.SignalEntry` to designate the entry the synchronized portion of the plan and the use of `watcher.Barrier` to wait for others to reach the same state. In this test case, the first plan to create a writer will act as th coordinator and all the others will follow.

```go
func ExampleSync(runenv *runtime.RunEnv) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	watcher, writer := sync.MustWatcherWriter(ctx, runenv)
	defer watcher.Close()
	defer writer.Close()

	runenv.RecordMessage("Waiting for network initialization")
	if err := sync.WaitNetworkInitialized(ctx, runenv, watcher); err != nil {
		return err
	}
	runenv.RecordMessage("Network initilization complete")

	st := sync.Subtree{
		GroupKey:    "messages",
		PayloadType: reflect.TypeOf(""),
		KeyFunc: func(val interface{}) string {
			return val.(string)
		}}

	seq, err := writer.Write(ctx, &st, runenv.TestRun)
	if err != nil {
		return err
	}

	runenv.RecordMessage("My sequence ID: %d", seq)

	readyState := sync.State("ready")
	startState := sync.State("start")

	if seq == 1 {
		runenv.RecordMessage("I'm the boss.")
		numFollowers := runenv.TestInstanceCount - 1
		runenv.RecordMessage("Waiting for %d instances to become ready", numFollowers)
		err := <-watcher.Barrier(ctx, readyState, int64(numFollowers))
		if err != nil {
			return err
		}
		runenv.RecordMessage("The followers are all ready")
		runenv.RecordMessage("Ready...")
		time.Sleep(1 * time.Second)
		runenv.RecordMessage("Set...")
		time.Sleep(5 * time.Second)
		runenv.RecordMessage("Go!")
		_, err = writer.SignalEntry(ctx, startState)
		return err
	} else {

		rand.Seed(time.Now().UnixNano())
		sleepTime := rand.Intn(10)
		runenv.RecordMessage("I'm a follower. Signaling ready after %d seconds", sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)
		_, err = writer.SignalEntry(ctx, readyState)
		if err != nil {
			return err
		}
		err = <-watcher.Barrier(ctx, startState, 1)
		if err != nil {
			return err
		}
		runenv.RecordMessage("Received Start")
		return nil
	}
}

```

When we run this, we need to add more than one instance. For testing, lets run it with four instances and have a look at the output. Using `--collect` we can get the output from each instance individually in json log format. In my run, instance 0 was the leader and the others were followers, each doing work \(sleeping\) that takes a different amount of time. 

{% tabs %}
{% tab title="Leader" %}
{% code title="" %}
```bash
unzip 81bd67649bae.zip  # This was the name of my outputs file

jq '.event.message' <81bd67649bae/single/0/run.out
"registering default http handler at: http://[::]:6060/ (pprof: http://[::]:6060/debug/pprof/)"
null
"waiting for pushgateway to become accessible"
"pushgateway is up at localhost:9091; pushing metrics every 5s."
"Waiting for network initialization"
"network initialisation successful"
"Network initilization complete"
"My sequence ID: 1"
"I'm the boss."
"Waiting for 3 instances to become ready"
"The followers are all ready"
"Ready..."
"Set..."
"Go!"
null
"test run completed. waiting a few seconds for the metrics scraper."
"io closed"
"metrics done"

```
{% endcode %}
{% endtab %}

{% tab title="Follower1" %}
```
jq '.event.message' <81bd67649bae/single/1/run.out

"registering default http handler at: http://[::]:46409/ (pprof: http://[::]:46409/debug/pprof/)"
null
"waiting for pushgateway to become accessible"
"pushgateway is up at localhost:9091; pushing metrics every 5s."
"Waiting for network initialization"
"network initialisation successful"
"Network initilization complete"
"My sequence ID: 3"
"I'm a follower. Signaling ready after 5 seconds"
"Received Start"
null
"test run completed. waiting a few seconds for the metrics scraper."
"io closed"
"metrics done"

```
{% endtab %}

{% tab title="Follower2" %}
```
jq '.event.message' <81bd67649bae/single/2/run.out

"registering default http handler at: http://[::]:36455/ (pprof: http://[::]:36455/debug/pprof/)"
null
"waiting for pushgateway to become accessible"
"pushgateway is up at localhost:9091; pushing metrics every 5s."
"Waiting for network initialization"
"network initialisation successful"
"Network initilization complete"
"My sequence ID: 2"
"I'm a follower. Signaling ready after 7 seconds"
"Received Start"
null
"test run completed. waiting a few seconds for the metrics scraper."
"io closed"
"metrics done"

```
{% endtab %}

{% tab title="Follower3" %}
```
jq '.event.message' <81bd67649bae/single/3/run.out

"registering default http handler at: http://[::]:34245/ (pprof: http://[::]:34245/debug/pprof/)"
null
"waiting for pushgateway to become accessible"
"pushgateway is up at localhost:9091; pushing metrics every 5s."
"Waiting for network initialization"
"network initialisation successful"
"Network initilization complete"
"My sequence ID: 4"
"I'm a follower. Signaling ready after 2 seconds"
"Received Start"
null
"test run completed. waiting a few seconds for the metrics scraper."
"io closed"
"metrics done"

```
{% endtab %}
{% endtabs %}

Synchronization bariers are curcial for many complex test plans and for any test plan which has distinct initialization and test phases, this is a crucial component. 

