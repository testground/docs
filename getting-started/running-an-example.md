# Running an example testplan

The first testplan that we will run is the `network` testplan and the `ping-pong` testcase.

The `ping-pong` testcase starts 2 testplan instances: one that listens on a TCP socket and another that dials it. The testcase exercises the synchronization service as well as the sidecar traffic shaping and IP allocation functionality.

Configure `$TESTGROUND_HOME` to a directory that would hold your testplan:

```
$ mkdir -p ~/.testground

$ export TESTGROUND_HOME=~/.testground
```

Copy the `network` testplan into the `$TESTGROUND_HOME` directory.

```
$ mkdir -p $TESTGROUND_HOME/plans

$ cd $GOPATH/src/github.com/testground/testground

$ cp -r plans/network $TESTGROUND_HOME/plans
```

Run the `network`testplan and the `ping-pong` testcase with the `docker:go` builder and the `local:docker` runner.

```
$ testground run single \
                 --plan=network \
                 --testcase=ping-pong \
                 --builder=docker:go \
                 --runner=local:docker \
                 --instances=2
```

After the testplan run execution concludes, you should see the following message:

```
INFO run finished successfully {"req_id": "d570c53a", "plan": "network", "case": "ping-pong", "runner": "local:docker", "instances": 2}

INFO finished run with ID: 5222e5df793b
```

All testplan run outputs and logs are stored at `$TESTGROUND_HOME/data/outputs/local_docker/network`

You could also fetch them with the following command:

```
$ testground collect --runner=local:docker 5222e5df793b
```

