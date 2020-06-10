# Client-Server communication

## Build <a id="build"></a>

`POST /build`

Build sends a mime/multipart \`build\` request to the daemon. A build request comprises the following parts:

* Part 1 \(Content-Type: application/json\): composition json \(see below\)
* Part 2 \(Content-Type: application/zip\): test plan source
* Part 3 \(optional, Content-Type: application/zip\): linked sdk

The request is the same kind multipart mime message archive your email administrator knows and loves so well. With this archive in hand, the plan is POST'ed the testground daemon.

## Run <a id="run"></a>

`POST /run`

A Run request consists of a composition json \(see below\)

​

## Composition JSON <a id="composition-json"></a>

```text
{  "composition": {    "metadata": {      "name": string,      "author": string,    },    "global": {      "plan": string,      "case": string,      "total_instances": int,      "builder": string,      "build_config": object,      "runner": string,      "run_config": string,    },    "groups": [      "id": string,      "resources": {        "memory": string,        "cpu": string,      },      "instances": {        "count": int,        "percentage": float,      },      "build":      "run":    ],  },}
```

## Collect outputs <a id="collect-outputs"></a>

`POST /outputs`

```text
{  "runner": string,  "run_id": string,}
```

## Terminate Jobs <a id="terminate-jobs"></a>

`POST /terminate`

## Healthcheck <a id="healthcheck"></a>

`POST /healthcheck`

```text
{  "runner": string,  "fix": bool,}
```

