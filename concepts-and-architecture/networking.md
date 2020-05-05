# Networking

All Testground runners _except_ for the `local:exec` runner have two networks: **a `control` and a `data` network.**

* Test plan instances communicate with each other over the **`data`** network.
* Test plan instances communicate with infrastructure services such as the [sync service](sync-service.md) and **InfluxDB** over the control network.

{% hint style="info" %}
This separation allows test instances to simulate disconnected scenarios; intermittent, failing connectivity; or high-latency scenarios without affecting other infrastructural comms.
{% endhint %}

{% hint style="warning" %}
The `local:exec` runner will use your machine's local network interfaces. **For now, this runner doesn't support traffic shaping.**
{% endhint %}

### Control Network

The `control` network is used to be communicate with Testground services, such as the [sync service](sync-service.md) or **InfluxDB**. You don't need to do anything special to configure or use this network: the [sidecar](sidecar.md) will do it for you automatically.

After the [sidecar](sidecar.md) is finished _initializing the network_, it should be impossible to use the `control` network to communicate with other test plan instances.

However, a good test plan should avoid listening on and/or announcing the `control` network _anyways_ to ensure that it doesn't interfere with the test. **Your test plan should always communicate via the data network; continue reading.**

### Data Network

The `data` network, used for communication between test plan instances, will be assigned a `B` block in the IP range `16.0.0.0/4`. Given the B block `X.Y.0.0`, `X.Y.0.1` is always the gateway and shouldn't be used by the test.

The subnet used will be passed to the test instance via the runtime environment \(as `TestSubnet`\).

{% hint style="success" %}
From the Go SDK, you can use the [`network.GetDataNetworkIP()`](https://pkg.go.dev/github.com/testground/sdk-go@v0.2.1/network?tab=doc#Client.GetDataNetworkIP) function to acquire your data network IP address.
{% endhint %}

