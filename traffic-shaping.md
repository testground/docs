# Traffic shaping

In Testground, a test instance can configure its own network \(IP address, jitter, latency, bandwidth, etc.\) using a network client. See the [network package of the Go SDK](https://pkg.go.dev/github.com/testground/sdk-go@v0.2.1/network?tab=doc) for more information.

## Imports

For this example, you'll need the following packages:

```go
import (
    "net"

    "github.com/testground/sdk-go/network"
    "github.com/testground/sdk-go/runtime"
    "github.com/testground/sdk-go/sync"
)
```

## Pre-check and preparation

First, check to make sure the sidecar is available. At the moment, it's only available on docker-based runners. If it's not available, just skip any networking config code and return.

```go
// runenv is the test instances run environment (runtime.RunEnv).
if !runenv.TestSidecar {
    return
}

// client is the *sync.Client
netclient := network.NewClient(client, runenv)
```

## Network initialization

Wait for the sidecar to initialize the network for this test plan instance. See the [Networking](concepts-and-architecture/networking.md) section for more details.

```go
netclient.MustWaitNetworkInitialized(ctx)
```

If you don't want to customize the network \(set IP addresses, latency, etc.\), you can stop here.

## Configure traffic shaping

Once the network is ready, you'll need to actually _configure_ your network. To "shape" traffic, set the `Default` `LinkShape`. You can use this to set latency, bandwidth, jitter, etc.

```go
config := network.Config{
    // Control the "default" network. At the moment, this is the only network.
    Network: "default",

    // Enable this network. Setting this to false will disconnect this test
    // instance from this network. You probably don't want to do that.
    Enable:  true,

    // Set the traffic shaping characteristics.
    Default: network.LinkShape{
        Latency:   100 * time.Millisecond,
        Bandwidth: 1 << 20, // 1Mib
    },

    // Set what state the sidecar should signal back to you when it's done.
    CallbackState: "network-configured",
}
```

{% hint style="info" %}
This sets _egress_ \(outbound\) properties on the link. These settings must be symmetric \(applied on both sides of the connection\) to work properly \(unless asymmetric bandwidth/latency/etc. is desired\).
{% endhint %}

{% hint style="info" %}
Per-subnet traffic shaping is a desired but unimplemented feature. The sidecar will reject configs with per-subnet rules set in `network.Config.Rules`.
{% endhint %}

## **\(Optional\) Changing your IP address**

If you don't specify an IPv4 address when configuring your network, your test instance will keep the default assignment. However, if desired, a test instance can change its IP address at any time.

First, you'll need some kind of unique sequence number to ensure you don't pick conflicting IP addresses. If you don't already have some form of unique sequence number at this point in your tests, use the sync service to get one:

```go
topic := sync.NewTopic("ip-allocation", "")
seq := sync.MustPublish(ctx, topic, "")
```

Once you have a sequence number, you can set your IP address from one of the available subnets:

```go
// copy the test subnet.
config.IPv4 = &*runenv.TestSubnet
// Use the sequence number to fill in the last two octets.
//
// NOTE: Be careful not to modify the IP from `runenv.TestSubnet`.
// That could trigger undefined behavior.
ipC := byte((seq >> 8) + 1)
ipD := byte(seq)
config.IPv4.IP = append(config.IPv4.IP[0:2:2], ipC, ipD)
```

You cannot currently set an IPv6 address.

## Apply the configuration

Applying the network configuration will post the configuration to the sync service, from where the appropriate instance of sidecar will consume it to apply the rules via Netlink. Once it is done, it will signal back on the `CallbackState`.

{% hint style="info" %}
The network API will, by default, wait for `runenv.TestInstanceCount` instances to have signalled on the `CallbackState`. If you want to wait for a different number of instances, such as if only a subset of instances actually apply traffic shaping rules, you can set the `CallbackTarget` value in the configuration.
{% endhint %}

```go
err := netclient.ConfigureNetwork(ctx, config)
if err != nil {
    runenv.Abort(err)
    return
}
```

## Appendix: What the sidecar does

1. The sidecar reads the network configuration from the sync service.
2. The sidecar applies network configurations requested by test plan instances.
3. The sidecar signals the configured `CallbackState` when done.

