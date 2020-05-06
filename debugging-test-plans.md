# Debugging test plans

While writing a test plan, there are a few ways to troubleshoot. This document explains a few options available for finding bugs in test plans and troubleshooting failures. On this page, I will introduce bugs intentionally so we can see how the system behaves and troubleshoot it.

## Build errors

```text
$ testground plan create --plan planbuggy
```

The command above will create a default `quickstart` test cases. Unfortunately, for our purposes, the plan has no bugs. Edit `main.go` so it contains the following buggier code.

{% code title="main.go" %}
```go
package main

import "github.com/testground/sdk-go/runtime"

func main() {
	runtime.Invoke(run)
}

func run(runenv *runtime.RunEnv) error {
	// No closing quote, will not build.
	runenv.RecordMessage("Hello Bugs)
	return nil
}
```
{% endcode %}

### How it looks in terminal output

When this plan runs, the code is sent to the daemon to be built. Of course, this will fail. Notice that the output comes in several sections.  The section labeled `Server output` shows us the error encountered by our builder.

```bash
$ testground run single --plan planbuggy --testcase quickstart --runner local:exec --builder exec:go --instances 1

May  5 00:31:15.650020	INFO	using home directory: /home/cory/testground
May  5 00:31:15.650143	INFO	no .env.toml found at /home/cory/testground/.env.toml; running with defaults
May  5 00:31:15.650182	INFO	testground client initialized	{"addr": "localhost:8042"}
May  5 00:31:15.651180	INFO	using home directory: /home/cory/testground
May  5 00:31:15.651300	INFO	no .env.toml found at /home/cory/testground/.env.toml; running with defaults
May  5 00:31:15.651339	INFO	testground client initialized	{"addr": "localhost:8042"}
May  5 00:31:15.651772	INFO	test plan source at: /home/cory/testground/plans/planbuggy

>>> Server output:

May  5 00:31:15.662268	INFO	performing build for groups	{"req_id": "be8cc8ee", "plan": "planbuggy", "groups": ["single"], "builder": "exec:go"}
May  5 00:31:15.906353	ERROR	go build failed: # github.com/coryschwartz/planbuggy
./main.go:10:35: newline in string
./main.go:10:35: syntax error: unexpected newline, expecting comma or )
./main.go:11:2: syntax error: unexpected return at end of statement
	{"req_id": "be8cc8ee"}
May  5 00:31:15.906505	INFO	build failed	{"req_id": "be8cc8ee", "plan": "planbuggy", "groups": ["single"], "builder": "exec:go", "error": "failed to run the build; exit status 2"}
May  5 00:31:15.906563	WARN	engine build error: failed to run the build; exit status 2	{"req_id": "be8cc8ee"}

>>> Error:

engine build error: failed to run the build; exit status 2
```

In this case, the error is pretty straightforward, but in a more complex plan, this output can be difficult to parse. So what can you do?

### Using standard debugging tools

Test plans are regular executables which accept configuration through environment variables. Because of this, you can test by compiling, testing, and running code. Except for the sync service, the code can be tested outside of Testground. Let's test the code this time without sending it to the Testground daemon. Let's see what the same code looks like testing locally.

```bash
$ go test

./main.go:10:35: newline in string
./main.go:10:35: syntax error: unexpected newline, expecting comma or )
./main.go:11:2: syntax error: unexpected return at end of statement
```

{% hint style="info" %}
If your plan relies on knowledge of the test plan or test case, this can be passed as an environment variable.
{% endhint %}

Now that output is much more readable!

I can't claim that build errors will always be as easy to diagnose as this one, but this feature enables plan writers to employ traditional debugging techniques or other debugging tools which they are already familiar.

## Debugging with message output

The next technique is useful for plans which build correctly and you want to observe the behaviour for debugging. If you have ever debugged a program by adding logging or printing to the screen, you know exactly what I'm talking about. On Testground plans can emit events and messages.

```text
runenv.RecordEvent("this is a message")
```

 Another thing which might be useful for debugging is events.  Just like messages, events can be used as a point-in-time caputre of the current state. Events are included in the outputs collection. They are recorded in the order they occur for each plan instance. We created R\(\) and D\(\) metrics collectors \(results and debugging\).  The difference between these two is that debugging is sent to the metrics pipeline fairly quickly whereas results are collected at the end of a test run.

```go
var things int
for {
  // work work work...
  things++
  runenv.D().RecordPoint("how_many_things", things)
}
runenv.R().RecordPoint("total_things", things)
```

To see how this works, let's use [ron swanson's classic dilemma](http://adit.io/posts/2013-05-11-The-Dining-Philosophers-Problem-With-Ron-Swanson.html).

In the following plan, five ~~philosophers~~ Ron Swansons sit at a table with five forks between them. Unfortunately, there is an implementation bug and these Ron Swansons will be be here forever. Add some debugging messages using `runenv.RecordMessage` to see if you can straighten this whole thing out \(hint: answer is in the second tab\)

{% tabs %}
{% tab title="exercise" %}
```go
package main

import (
	"github.com/testground/sdk-go/runtime"
	"sync"
)

type Fork struct {
	id int
	m  *sync.Mutex
}

type Swanson struct {
	id                  int
	meals               int
	leftFork, rightFork *Fork
	wg                  *sync.WaitGroup
}

func (ron *Swanson) Feast(runenv *runtime.RunEnv) {
	ron.leftFork.m.Lock()
	ron.rightFork.m.Lock()
	r.D().RecordPoint("eats", 1)
	
	ron.leftFork.m.Unlock()

	ron.meals = ron.meals - 1
	if ron.meals > 0 {
		runenv.Message("Still hungry. %d more meals to fill me up.", ron.meals)
		ron.Feast(runenv)
	} else {
		runenv.Message("All done. For now...")
		ron.wg.Done()
	}
}

func main() {
	runtime.Invoke(run)
}

func run(runenv *runtime.RunEnv) error {
  // Five hungry rons eat 10 plates of food each.
	countRons := 5
	countMeals := 10
	wg := sync.WaitGroup{}

	// Create forks
	forks := make([]*Fork, countRons)
	for i := 0; i < countRons; i++ {
		forks[i] = &Fork{
			id: i,
			m:  new(sync.Mutex),
		}
	}

	// Each ron swanson has has a fork to his left and right
	rons := make([]*Swanson, countRons)
	wg.Add(countRons)
	for i := 0; i <= countRons; i++ {
		rons[i] = &Swanson{
			id:        i,
			leftFork:  forks[i%countRons],
			rightFork: forks[(i+1)%countRons],
			meals:     countMeals,
			wg:        &wg,
		}
		go rons[i].Feast(runenv)
	}

	wg.Wait()
	runenv.RecordMessage("all rons have eaten")
	return nil
}

```
{% endtab %}

{% tab title="Solution" %}
```
Line 24.
Ron puts down his left fork, but forgets to put down his right fork!
Add another line to unlick the rigtFork mutex to fix this problem.

Line 58.
We think there are an equal number of Rons and forks, but we are off
by one. Two Rons are sharing the same seat! this will not do.
```
{% endtab %}
{% endtabs %}

If you can successfully debug this code, you will see each ron finish his meals and finally the  end message "**all rons have eaten**"

### Collecting outputs vs viewing messages in the terminal

When using the local runners, with a relatively small number of plan instances it is fairly easy to view outputs in the terminal runner. I recommend troubleshooting the plan with a small number of instances. The same messages you can see in your terminal are also available in outputs collections.

For more information about this, see [Analyzing the results](https://app.gitbook.com/@protocol-labs/s/testground/~/drafts/-M6X2x7PG-JL0LAa-bnw/analyzing-the-results).

## Accessing profile data

All Go test plans have profiling enabled by default.

For information about using Go's `pprof` and generating graphs and reports, I recommend you [start here](https://golang.org/pkg/net/http/pprof/).

On Testground gaining access to the `pprof` port can sometimes be non-obvious. Allow me to explain how to get access to port 6060 on each of our runners:

{% tabs %}
{% tab title="local:exec" %}
```bash
# Plans in the local:exec runner operate in the default namespace.
# the first plan will grab port 6060 and the any additional will listen
# on a random port.
#
# look for the following messages in the test run output to figure out
# the URL where to access the pprof endpoint of each instance:
#
# May  6 14:32:10.146239	INFO	0.0174s    MESSAGE << instance   1 >>
# registering default http handler at: http://[::]:6060/
# (pprof: http://[::]:6060/debug/pprof/)
#
# May  6 14:32:10.146535	INFO	0.0179s    MESSAGE << instance   2 >>
# registering default http handler at: http://[::]:64912/
# (pprof: http://[::]:64912/debug/pprof/)


$ echo "top" | go tool pprof http://localhost:6060/debug/pprof/heap

Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
top: open top: no such file or directory
Fetched 1 source profiles out of 2
Saved profile in /home/cory/pprof/pprof.exec-go--planbuggy-ba70706f7cd1.alloc_objects.alloc_space.inuse_objects.inuse_space.003.pb.gz
File: exec-go--planbuggy-ba70706f7cd1
Type: inuse_space
Time: May 5, 2020 at 1:05am (PDT)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) Showing nodes accounting for 544.67kB, 100% of 544.67kB total
      flat  flat%   sum%        cum   cum%
  544.67kB   100%   100%   544.67kB   100%  net.open
         0     0%   100%   544.67kB   100%  internal/singleflight.(*Group).doCall
         0     0%   100%   544.67kB   100%  net.(*Resolver).goLookupIPCNAMEOrder
         0     0%   100%   544.67kB   100%  net.(*Resolver).lookupIP
         0     0%   100%   544.67kB   100%  net.(*Resolver).lookupIPAddr.func1
         0     0%   100%   544.67kB   100%  net.glob..func1
         0     0%   100%   544.67kB   100%  net.goLookupIPFiles
         0     0%   100%   544.67kB   100%  net.lookupStaticHost
         0     0%   100%   544.67kB   100%  net.readHosts

```
{% endtab %}

{% tab title="local:docker" %}
```bash
# local:docker runner has isolated networking for each container.
# An ephemeral port is assigned in the host network namespace which maps
# to port 6060 in the containers namespace.

# Find out the port assignment using `docker ps` and then profile the
# appropriate port. For example, in this case I have three plans running
# in different containers. To profile one of these, I can point pprof tool
# to a port in the host namespace.

$ docker ps -f 'name=tg-' --format "{{.Ports}}"
0.0.0.0:33279->6060/tcp
0.0.0.0:33278->6060/tcp
0.0.0.0:33280->6060/tcp


$ echo top | go tool pprof http://localhost:33279/debug/pprof/heap 

Fetching profile over HTTP from http://localhost:33279/debug/pprof/heap
Saved profile in /home/cory/pprof/pprof.testplan.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
File: testplan
Type: inuse_space
Time: May 5, 2020 at 1:15am (PDT)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) Showing nodes accounting for 1026.44kB, 100% of 1026.44kB total
Showing top 10 nodes out of 11
      flat  flat%   sum%        cum   cum%
     514kB 50.08% 50.08%      514kB 50.08%  bufio.NewWriterSize (inline)
  512.44kB 49.92%   100%   512.44kB 49.92%  runtime.allocm
         0     0%   100%      514kB 50.08%  net/http.(*conn).serve
         0     0%   100%      514kB 50.08%  net/http.newBufioWriterSize
         0     0%   100%   512.44kB 49.92%  runtime.handoffp
         0     0%   100%   512.44kB 49.92%  runtime.mcall
         0     0%   100%   512.44kB 49.92%  runtime.newm
         0     0%   100%   512.44kB 49.92%  runtime.park_m
         0     0%   100%   512.44kB 49.92%  runtime.schedule
         0     0%   100%   512.44kB 49.92%  runtime.startm

```
{% endtab %}

{% tab title="cluster:k8s" %}
```bash
# When using the Kubernetes cluster:k8s runner, each container runs in a
# separate pod. Access the profiling port throught the Kubernetes API using
# a port forward or kubectl proxy. The following example sets up a port
# forward to a sidecar pod and then runs a profiler against it.

$ kubectl port-forward testground-sidecar 6060:6060

$ echo top | go tool pprof http://localhost:6060/debug/pprof/heap

Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
Saved profile in /home/cory/pprof/pprof.testground.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
File: testground
Type: inuse_space
Time: May 5, 2020 at 1:41am (PDT)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) Showing nodes accounting for 3112.91kB, 100% of 3112.91kB total
Showing top 10 nodes out of 21
      flat  flat%   sum%        cum   cum%
 1025.94kB 32.96% 32.96%  1025.94kB 32.96%  regexp/syntax.(*compiler).inst (inline)
  548.84kB 17.63% 50.59%   548.84kB 17.63%  github.com/markbates/pkger/internal/takeon/github.com/markbates/hepa/filters.glob..func3
     514kB 16.51% 67.10%      514kB 16.51%  regexp.mergeRuneSets.func2 (inline)
  512.08kB 16.45% 83.55%  2052.02kB 65.92%  regexp.compile
  512.05kB 16.45%   100%   512.05kB 16.45%  runtime.systemstack
         0     0%   100%   512.62kB 16.47%  github.com/dustin/go-humanize.init.0
         0     0%   100%  1026.08kB 32.96%  github.com/go-playground/validator/v10.init
         0     0%   100%   548.84kB 17.63%  github.com/markbates/pkger/internal/takeon/github.com/markbates/hepa/filters.init
         0     0%   100%   513.31kB 16.49%  k8s.io/apimachinery/pkg/util/naming.init
         0     0%   100%  2052.02kB 65.92%  regexp.Compile (inline)

```
{% endtab %}
{% endtabs %}



