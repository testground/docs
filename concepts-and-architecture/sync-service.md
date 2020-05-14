---
description: >-
  The synchronization service offers simple, but powerful and composable
  primitives to coordinate and choreograph distributed test workloads.
---

# Synchronization service

## Context

In distributed testing workloads, instances need to perform actions around a scripted series of steps. We are essentially modelling a distributed state machine, spanning all workloads that participate in the test run.

**To make this possible, test instances need to synchronize and coordinate around a choreography.**

Some concrete coordination problems that may emerge in a 1000-instance run of, say, a peer-to-peer network, are:

1. **Role assignment.** How do we determine which instances are bootstrappers, service providers, service consumers, seeds, leeches?
2. **Signalling** that instances have arrived at a particular point in the test plan, and are now ready to advance to the next stage, once N instances are ready too.
3. **Communicating out-of-band information,** such as endpoint addresses, identities, random content identifiers, secrets, etc.

## Possible solutions

There are various ways of implementing such coordination. We could either adopt:

1. **❌ a command-and-control model:** by having test plans deploy a centralized coordinator instance that acts like a "conductor", telling each child instance participating in the run what it needs to do next.
   * this model performs poorly in terms of resiliency.
   * this model introduces a scheduling dependency: we need to deploy the coordinator first, obtain its address, and somehow communicate it to the children.
   * this model is complex in terms of design and development: test plan writers need to write the code that will run on the coordinator, as well as the state corresponding checkpoints in the children where an interaction with the coordinator must happen.
2. **✅ a distributed coordination model:** by coordinating test instances in a decentralized fashion. The same test plan runs on all machines, using an API that hits a common high-performance synchronization store, and offers distributed synchronization primitives like barriers, signals, pubsub, latches, semaphores, etc.

## Testground sync API

In Testground, test plans hit the APIs of the components under test directly, and whenever they need to synchronize with other participants in the run, they call out to **the sync API.**

Our **sync API** is extensively inspired by Apache ZooKeeper and the Apache Curator recipes, but we have chosen a non-durable, memory-only Redis instance for simplicity and performance reasons. Therefore, the **sync API** is a lightweight wrapper around a Redis client. All Testground runners deploy a Redis instance \(as a Docker container, or as a k8s pod\), and inject the address into test workloads.

The **sync API** offers simple, but powerful and composable synchronization primitives to coordinate and choreograph distributed test workloads. 

We have implemented the following primitives so far, and more are to come, such as locks, semaphores, leader election, etc. Take a look at the [Apache Curator recipes](https://curator.apache.org/curator-recipes/index.html) and the [Redisson project](https://github.com/redisson/redisson/wiki/8.-distributed-locks-and-synchronizers) to understand where our thinking is.

**Supported synchronization primitives**

* **State signalling and barriers:** instances can signal **entry into states**, and can **await** until N instances enter that state. Example use cases: wait until all instances have started, wait until instances in group "adders" have added a file to their repo, wait until all nodes have bootstrapped, etc.

```go
bootstrapped := sync.State("bootstrapped")

// once they've bootstrapped, instances signal on the "bootstrapped" state
seq := client.MustSignalEntry(bootstrapped)

fmt.Printf("I am instance number %d entering the 'bootstrapped' state\n", seq)

// then they wait for all other instances to have bootstrapped
client.MustBarrier(bootstrapped, runenv.TestInstanceCount)

fmt.Println("All instances have entered the 'bootstrapped' state")
```

* **Topic publishing and subscribing:** instances can emit arbitrary data on topics, and can subscribe to consume that data. Example use cases: communicating endpoint addresses, communicating unique identifiers, etc.

```go
// topic 'addresses', of type string (we infer the type from the 2nd arg)
addresses := sync.NewTopic("addresses", "")

// listen on an random port
listener, _ := net.Listen("tcp", "0.0.0.0:0")

// instances publish their endpoint addresses on the 'addresses' topic
seq := client.MustPublish(ctx, addresses, listener.Addr().String())

fmt.Printf("I am instance number %d publishing to the 'addresses' topic\n", seq)

// consume all addresses from all peers
ch := make(chan string)
sub := client.MustSubscribe(ctx, topic, ch)

for i := 0; i < runenv.TestInstanceCount; i++ {
    addr := <-ch
    fmt.Println("received addr:", addr)
}

// we cancel the context we passed to the subscription
cancel()
```

For more information, refer to the [godocs for the sync package of the Go SDK](https://pkg.go.dev/github.com/testground/sdk-go@v0.2.1/sync?tab=doc).

