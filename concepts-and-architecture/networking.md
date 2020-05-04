# Networking

All Testground runners _except_ for the `local:exec` runner have two networks: a `control` and a `data` network.

* Testplan instances communicate with each other over the `data` network.
* Testplan instances communicate with the `sync service` over the control network.

The `local:exec` runner will use your machine's local network interfaces. For now, this runner doesn't support traffic shaping.

### Control Network

The `control` network should only be used to communicate with Testground services, such as the `sync service`.

After the `sidecar` is finished `initializing the network`, it should be impossible to use the `control` network to communicate with other test plan instances. However, a good test plan should avoid listening on and/or announcing the `control` network _anyways_ to ensure that it doesn't interfere with the test.

### Data Network

The `data` network, used for all inter-testplan-instance communication, will be assigned a B block in the IP range `16.0.0.0/4`. Given the B block X.Y.0.0, X.Y.0.1 is always the gateway and shouldn't be used by the test.

The subnet used will be passed to the test instance via the runtime environment \(as `TestSubnet`\).

