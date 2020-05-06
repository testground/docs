# Communication between instances

For some test plans, it is useful to pass information from one instance to anther. In addition to direct network connectivity, test plans can pass information between instances using the testground sync service. In this tutorial, we will explore typed typed message passing through the testground sync service.

Under the covers, communication between instances happens through a redis server using `XADD` and `XREAD` to to stream json-encoded objects. The producer of the data writes to the stream, and the receiver subscribes to it.

Lets create a plan in which one of the plans produces a structure which is re-constructed on the distant end.  First, I'll show short snipits with the relevent informaiton, and the whole test plan will be shown at the end.

#### Setting up the subtree

This is the kind of object we will be transferring.

```go
type Transferrable struct {
	Name          string
	FavoriteSport int
	CareWhoKnows  bool
}
```

The object will be transferred over a subtree. Think of the subtree as a named and typed channel for transferring objects between plan instances. This subtree is named `transfer-key` and the type of object I expect to get out of it is `pointer to Transferrable`.

```go
	st := sync.Subtree{
		GroupKey:    "transfer-key",
		PayloadType: reflect.TypeOf(&Transferrable{}),
	}
```

#### Writing to the subtree

To write to the subtree, create a `sync.Writer` and use it to write to the subtree we have just defined. Notice that the type of the object we are writing must match the type of the subtree.

```go
	ctx := context.Background()

	writer, err := sync.NewWriter(ctx, runenv)
	if err != nil {
		return err
	}
	writer.Write(ctx, &st, &Transferrable{"Guy#1", 1, false})

```

#### Reading from the subtree

Subscribe to the subtree we created earlier by creating a watcher and setting up a channel to receive the received interfaces.

```go
	tch := make(chan *Transferrable)
	watcher, err := sync.NewWatcher(ctx, runenv)
	if err != nil {
		return err
	}
	watcher.Subscribe(ctx, &st, tch)
```

#### But who subscribes and who publishes?

This question is left up to the plan writer, and certainly different situations will call for different implementations. In this example, all the plans will publish and all will subscribe, but there are some scenarios where this is inappropriate.



#### Full Example

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"

	"github.com/ipfs/testground/sdk/runtime"
	"github.com/ipfs/testground/sdk/sync"
)

type Sport int

const (
	football Sport = iota
	tennis
	hockey
	golf
)

func (s Sport) String() string {
	return [...]string{"football", "basketball", "baseball", "hockey"}[s]
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
	runtime.Invoke(run)
}

func run(runenv *runtime.RunEnv) error {
	ctx := context.Background()
	watcher, writer := sync.MustWatcherWriter(ctx, runenv)
	st := sync.Subtree{
		GroupKey:    "transfer-key",
		PayloadType: reflect.TypeOf(&Transferrable{}),
	}

	// Configure the test
	myName := fmt.Sprintf("Guy#%d", rand.Int()%100)
	mySport := Sport(rand.Int() % 4)
	howMany := runenv.TestInstanceCount

	// Publisher
	writer.Write(ctx, &st, &Transferrable{myName, mySport, false})

	// Wait until published
	writer.SignalEntry(ctx, "ready")
	<-watcher.Barrier(ctx, "ready", int64(howMany))

	// Subscriber
	tch := make(chan *Transferrable)
	watcher.Subscribe(ctx, &st, tch)

	for i := 0; i < howMany; i++ {
		t := <-tch
		runenv.RecordMessage("%s", t)
	}

	return nil
}
```

Run with multiple instances, like this: 

```text
$ testground run single -p quickstart -t quickstart -b exec:go -r local:exec -i 2
```

{% hint style="info" %}
Notice that instances is set to 2. Two instances will run at the same time.
{% endhint %}

