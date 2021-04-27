# Communication between instances

For some test plans, it is useful to pass information from one instance to another. In addition to direct network connectivity, test plans can pass information between instances using the Testground `sync service`.

In this tutorial, we will explore typed message passing through the Testground sync service.

Lets create a plan in which one of the plans produces a `struct` which is re-constructed on the distant end. First, I'll show short snippets with the relevant information, and the whole test plan will be shown at the end.

### Setting up the`topic`

`Transferrable` is the value type we will be transferring.

```go
type Transferrable struct {
	Name          string
	FavoriteSport int
	CareWhoKnows  bool
}
```

The value will be transferred over a `topic`. Think of the `topic` as a named and typed channel for transferring values between plan instances. This topic is named `transfer-key` and the value type I expect to get out of it is `pointer to Transferrable`.

```go
st := sync.NewTopic("transfer-key", &Transferrable{})
```

### Publishing to a `topic`

To write to a topic, create a bounded client and use it to publish to the topic we have just defined.

```go
	ctx := context.Background()

	client := sync.MustBoundClient(ctx, runenv)
	defer client.Close()

	client.Publish(ctx, st, &Transferrable{"Guy#1", 1, false})
```

### Reading from a `topic`

Subscribe to the topic we created earlier and set up a channel to receive the values.

```go
	tch := make(chan *Transferrable)

	_, err = client.Subscribe(ctx, st, tch)
	if err != nil {
		panic(err)
	}
```

### Who subscribes and who publishes?

This question is left up to the plan writer, and certainly different situations will call for different implementations. In this example, all the plans will publish and all will subscribe, but there are scenarios where this is inappropriate.

### Full Example

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
	"github.com/testground/sdk-go/sync"
)

type Sport int

const (
	football Sport = iota
	tennis
	hockey
	golf
)

func (s Sport) String() string {
	return [...]string{"football", "tennis", "hockey", "golf"}[s]
}

type Transferrable struct {
	Name          string
	FavoriteSport Sport
	CareWhoKnows  bool
}

func (t *Transferrable) String() string {
	msg := fmt.Sprintf("%s: I like %s", t.Name, t.FavoriteSport)
	if t.CareWhoKnows {
		return msg + " and I really care!"
	}
	return msg + " and I don't care who knows!"
}

func main() {
	run.Invoke(runf)
}

func runf(runenv *runtime.RunEnv) error {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()
	client := sync.MustBoundClient(ctx, runenv)
	defer client.Close()

	st := sync.NewTopic("transfer-key", &Transferrable{})

	// Configure the test
	myName := fmt.Sprintf("Guy#%d", rand.Int()%100)
	mySport := Sport(rand.Int() % 4)
	howMany := runenv.TestInstanceCount

	// Publish my entry
	client.Publish(ctx, st, &Transferrable{myName, mySport, false})

	// Wait until all instances have published entries
	readyState := sync.State("ready")
	client.MustSignalEntry(ctx, readyState)
	<-client.MustBarrier(ctx, readyState, howMany).C

	// Subscribe to the `transfer-key` topic
	tch := make(chan *Transferrable)
	client.Subscribe(ctx, st, tch)

	for i := 0; i < howMany; i++ {
		t := <-tch
		runenv.RecordMessage("%s", t)
	}

	return nil
}
```

Run with multiple instances:

```text
$ testground run single -p quickstart -t quickstart -b exec:go -r local:exec -i 4
```

{% hint style="info" %}
Notice that instances is set to 4. Four instances will run at the same time.
{% endhint %}
