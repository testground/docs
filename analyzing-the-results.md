# Analyzing the results

## **Testground outputs structure**

Testground testplans emit `outputs` that you can collect and analyse after the testplan run concludes.

The outputs structure is as follows:

```text
/
|___ <group id> (coming from the composition)
|      |____ <instance id or index>
|                     |__ run.out
|                     |__ run.err
|                     |__ ...
|                     |__ dht_events.json
|                     |__ heap.out
|                     |__ cpuprofile01.out
|                     |__ cpuprofile02.out
|                     |__ cpuprofile03.out
|      |____ <instance id or index>
|      |____ <instance id or index>
|___ <group id> (coming from the composition)
|      |____ <instance id or index>
|      |____ <instance id or index>
```

## How to collect Testground outputs

The logic for `outputs` collection is runner-dependent, e.g.

* the `cluster:k8s` runner will download from AWS EFS.
* the `local:docker` runner will fetch the outputs from the `$TESTGROUND_HOME` directory.

```text
$ testground collect --runner=local:docker <run-id>

$ tar -xzvf <run-id>.tgz
```

Testground also supports automatic outputs collection after a testplan run completes, with the `--collect` flag, which instructs the Testground client to download the outputs in the current working directory after the run concludes.

```text
$ testground run single \
                 --plan=network \
                 --testcase=ping-pong \
                 --builder=docker:go \
                 --runner=local:docker \
                 --instances=2 \
                 --collect

...

>>> Result:

INFO    finished run with ID: 975b9bc15b3b

>>> Result:

INFO    created file: 975b9bc15b3b.tgz
```

## Processing the Testground outputs

At the moment Testground doesn't provide any tools for post-processing of outputs. You are free to submit ideas and pull requests and contribute to this feature.

One way to review all messages sent to `stdout` by testplan instances, is with the following command:

```text
$ cd <run-id>

$ find . | grep run.out | xargs awk '{print FILENAME, " >>> ", $0 }'
```

