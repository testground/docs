# Running an example test plan

The first test plan that we will run is the `network` test plan and the `ping-pong` test case.

The `ping-pong` test case starts 2 test plan instances: one that listens on a TCP socket and another that dials it. The test case exercises the synchronization service as well as the sidecar traffic shaping and IP allocation functionality.

Configure `$TESTGROUND_HOME` to a directory that would hold your test plan:

```
$ mkdir -p ~/.testground

$ export TESTGROUND_HOME=~/.testground
```

Copy the `network` test plan into the `$TESTGROUND_HOME` directory.

```
$ mkdir -p $TESTGROUND_HOME/plans

$ cd $GOPATH/src/github.com/testground/testground

$ cp -r plans/network $TESTGROUND_HOME/plans
```

Run the `network`testplan and the `ping-pong` test case with the `docker:go` builder and the `local:docker` runner.

```
$ testground run single \
                 --plan=network \
                 --testcase=ping-pong \
                 --builder=docker:go \
                 --runner=local:docker \
                 --instances=2
```

After the test plan run execution concludes, you should see the following message:

```
INFO run finished successfully {"req_id": "d570c53a", "plan": "network", "case": "ping-pong", "runner": "local:docker", "instances": 2}

INFO finished run with ID: 5222e5df793b
```

All test plan run outputs and logs are stored at `$TESTGROUND_HOME/data/outputs/local_docker/network`

You could also fetch them with the following command:

```
$ testground collect --runner=local:docker 5222e5df793b
```

