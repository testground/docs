# Keeping instances in sync

Sometimes individual test plan instances need to coordinate what they are doing. For this, we use [`barriers`](https://en.wikipedia.org/wiki/Barrier_%28computer_science%29).

The general concept is this -- as a plan reaches a phase of execution for which synchronisation is required, it will signal the other instances that it has reached that phase. Then, wait until a certain number of other instances reach the same state before continuing.

{% hint style="info" %}
Internally, this is implemented with a simple counter backed by a Redis database. To wait for a signal barrier means to wait until the counter reaches the specified number.
{% endhint %}

Let's have a look at a plan that waits for others! The plan displayed here is taken from the [example plans](https://github.com/ipfs/testground/tree/master/plans/example).

Notice the use of `client.MustSignalEntry(ctx, readyState)` to designate the entry to the synchronised portion of the plan and the use of `<-client.MustBarrier(ctx, readyState, numFollowers).C` to wait for others to reach the same state.

In this test case, the instance with `seq==1` will act as the coordinator and all the others will follow.

{% code title="plans/example/sync.go" %}
```go
func ExampleSync(runenv *runtime.RunEnv) error {
	var (
		readyState = sync.State("ready")
		startState = sync.State("start")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	client := sync.MustBoundClient(ctx, runenv)
	defer client.Close()

	netclient := network.NewClient(client, runenv)
	runenv.RecordMessage("Waiting for network initialization")

	netclient.MustWaitNetworkInitialized(ctx)
	runenv.RecordMessage("Network initilization complete")

	topic := sync.NewTopic("messages", "")

	seq, err := client.Publish(ctx, topic, runenv.TestRun)
	if err != nil {
		return err
	}

	runenv.RecordMessage("My sequence ID: %d", seq)

	if seq == 1 {
		runenv.RecordMessage("I'm the boss.")
		numFollowers := runenv.TestInstanceCount - 1

		runenv.RecordMessage("Waiting for %d instances to become ready", numFollowers)
		err := <-client.MustBarrier(ctx, readyState, numFollowers).C
		if err != nil {
			return err
		}

		runenv.RecordMessage("The followers are all ready")
		runenv.RecordMessage("Ready...")
		time.Sleep(1 * time.Second)
		runenv.RecordMessage("Set...")
		time.Sleep(5 * time.Second)
		runenv.RecordMessage("Go!")

		client.MustSignalEntry(ctx, startState)
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	sleepTime := rand.Intn(10)
	runenv.RecordMessage("I'm a follower. Signaling ready after %d seconds", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	client.MustSignalEntry(ctx, readyState)

	err = <-client.MustBarrier(ctx, startState, 1).C
	if err != nil {
		return err
	}

	runenv.RecordMessage("Received Start")
	return nil
}
```
{% endcode %}

When we run this, we need to add more than one instance. For testing, let's run it with four instances and have a look at the output. Using `--collect` we can get the output from each instance individually in JSON format.

```bash
# run the test plan
$ testground run single -p example -t sync -b exec:go -r local:exec -i 4 --collect

# extract the outputs archive
$ tar -xzvvf e217c3a8a089.tgz

$ cd e217c3a8a089

# print out the contents of all `run.out` from all instances
$ find . | grep run.out | xargs awk '{print FILENAME, " >>> ", $0 }' | sort

./single/0/run.out  >>>  {"ts":1588769459341224000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"registering default http handler at: http://[::]:51677/ (pprof: http://[::]:51677/debug/pprof/)"}}
./single/0/run.out  >>>  {"ts":1588769459341303000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"start","runenv":{"plan":"example","case":"sync","run":"e217c3a8a089","params":{},"instances":4,"outputs_path":"/Users/nonsense/tghome/data/outputs/local_exec/example/e217c3a8a089/single/0","network":"127.1.0.0/16","group":"single","group_instances":4}}}
./single/0/run.out  >>>  {"ts":1588769459424266000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Waiting for network initialization"}}
./single/0/run.out  >>>  {"ts":1588769459424312000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"network initialisation successful"}}
./single/0/run.out  >>>  {"ts":1588769459424330000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Network initilization complete"}}
./single/0/run.out  >>>  {"ts":1588769459425997000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"My sequence ID: 4"}}
./single/0/run.out  >>>  {"ts":1588769459426060000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"I'm a follower. Signaling ready after 1 seconds"}}
./single/0/run.out  >>>  {"ts":1588769460354489000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}
./single/0/run.out  >>>  {"ts":1588769465354614000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/0/run.out  >>>  {"ts":1588769470353134000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/0/run.out  >>>  {"ts":1588769470434227000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Received Start"}}
./single/0/run.out  >>>  {"ts":1588769470434337000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"finish","outcome":"ok"}}
./single/0/run.out  >>>  {"ts":1588769470443381000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"io closed"}}
./single/0/run.out  >>>  {"ts":1588769470458073000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}
./single/1/run.out  >>>  {"ts":1588769459341256000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"registering default http handler at: http://[::]:51679/ (pprof: http://[::]:51679/debug/pprof/)"}}
./single/1/run.out  >>>  {"ts":1588769459341329000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"start","runenv":{"plan":"example","case":"sync","run":"e217c3a8a089","params":{},"instances":4,"outputs_path":"/Users/nonsense/tghome/data/outputs/local_exec/example/e217c3a8a089/single/1","network":"127.1.0.0/16","group":"single","group_instances":4}}}
./single/1/run.out  >>>  {"ts":1588769459424352000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Waiting for network initialization"}}
./single/1/run.out  >>>  {"ts":1588769459424387000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"network initialisation successful"}}
./single/1/run.out  >>>  {"ts":1588769459424404000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Network initilization complete"}}
./single/1/run.out  >>>  {"ts":1588769459425996000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"My sequence ID: 3"}}
./single/1/run.out  >>>  {"ts":1588769459426058000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"I'm a follower. Signaling ready after 3 seconds"}}
./single/1/run.out  >>>  {"ts":1588769460354489000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}
./single/1/run.out  >>>  {"ts":1588769465354611000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/1/run.out  >>>  {"ts":1588769470351967000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/1/run.out  >>>  {"ts":1588769470431852000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Received Start"}}
./single/1/run.out  >>>  {"ts":1588769470432029000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"finish","outcome":"ok"}}
./single/1/run.out  >>>  {"ts":1588769470438221000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"io closed"}}
./single/1/run.out  >>>  {"ts":1588769470444752000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}
./single/2/run.out  >>>  {"ts":1588769459341164000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"registering default http handler at: http://[::]:6060/ (pprof: http://[::]:6060/debug/pprof/)"}}
./single/2/run.out  >>>  {"ts":1588769459341225000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"start","runenv":{"plan":"example","case":"sync","run":"e217c3a8a089","params":{},"instances":4,"outputs_path":"/Users/nonsense/tghome/data/outputs/local_exec/example/e217c3a8a089/single/2","network":"127.1.0.0/16","group":"single","group_instances":4}}}
./single/2/run.out  >>>  {"ts":1588769459423576000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Waiting for network initialization"}}
./single/2/run.out  >>>  {"ts":1588769459423623000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"network initialisation successful"}}
./single/2/run.out  >>>  {"ts":1588769459423634000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Network initilization complete"}}
./single/2/run.out  >>>  {"ts":1588769459425071000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"My sequence ID: 1"}}
./single/2/run.out  >>>  {"ts":1588769459425111000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"I'm the boss."}}
./single/2/run.out  >>>  {"ts":1588769459425128000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Waiting for 3 instances to become ready"}}
./single/2/run.out  >>>  {"ts":1588769460354489000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}
./single/2/run.out  >>>  {"ts":1588769463433338000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"The followers are all ready"}}
./single/2/run.out  >>>  {"ts":1588769463433419000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Ready..."}}
./single/2/run.out  >>>  {"ts":1588769464434839000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Set..."}}
./single/2/run.out  >>>  {"ts":1588769465354620000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/2/run.out  >>>  {"ts":1588769469438375000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Go!"}}
./single/2/run.out  >>>  {"ts":1588769469439899000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"finish","outcome":"ok"}}
./single/2/run.out  >>>  {"ts":1588769469445744000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"io closed"}}
./single/2/run.out  >>>  {"ts":1588769469448944000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 32 points"}}
./single/3/run.out  >>>  {"ts":1588769459341236000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"registering default http handler at: http://[::]:51678/ (pprof: http://[::]:51678/debug/pprof/)"}}
./single/3/run.out  >>>  {"ts":1588769459341314000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"start","runenv":{"plan":"example","case":"sync","run":"e217c3a8a089","params":{},"instances":4,"outputs_path":"/Users/nonsense/tghome/data/outputs/local_exec/example/e217c3a8a089/single/3","network":"127.1.0.0/16","group":"single","group_instances":4}}}
./single/3/run.out  >>>  {"ts":1588769459424359000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Waiting for network initialization"}}
./single/3/run.out  >>>  {"ts":1588769459424408000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"network initialisation successful"}}
./single/3/run.out  >>>  {"ts":1588769459424492000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Network initilization complete"}}
./single/3/run.out  >>>  {"ts":1588769459425979000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"My sequence ID: 2"}}
./single/3/run.out  >>>  {"ts":1588769459426057000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"I'm a follower. Signaling ready after 1 seconds"}}
./single/3/run.out  >>>  {"ts":1588769460354631000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}
./single/3/run.out  >>>  {"ts":1588769465352913000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/3/run.out  >>>  {"ts":1588769470351959000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 31 points"}}
./single/3/run.out  >>>  {"ts":1588769470434215000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"Received Start"}}
./single/3/run.out  >>>  {"ts":1588769470434336000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"finish","outcome":"ok"}}
./single/3/run.out  >>>  {"ts":1588769470443382000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"io closed"}}
./single/3/run.out  >>>  {"ts":1588769470446654000,"msg":"","group_id":"single","run_id":"e217c3a8a089","event":{"type":"message","message":"influxdb: uploaded 1 points"}}

```

