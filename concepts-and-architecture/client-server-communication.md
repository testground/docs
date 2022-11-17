# Client-Server communication

_This is an advanced section for those curious about the communication between client and server. This section is not needed to use Testground, but to understand the protocol between client and server._

## Build

`POST /build`

Build sends a mime/multipart \`build\` request to the daemon. A build request comprises the following parts:

* Part 1 \(Content-Type: application/json\): composition json \(see below\)
* Part 2 \(Content-Type: application/zip\): test plan source
* Part 3 \(optional, Content-Type: application/zip\): linked sdk

The request is the same kind multipart mime message archive your email administrator knows and loves so well. With this archive in hand, the plan is POST'ed the testground daemon.

## Run

`POST /run`

A Run request consists of a composition json \(see below\)

## Composition JSON

```javascript
{
  "composition": {
    "metadata": {
      "name": string,
      "author": string,
    },
    "global": {
      "plan": string,
      "case": string,
      "total_instances": int,
      "builder": string,
      "build_config": object,
      "runner": string,
      "run_config": string,
    },
    "groups": [
      "id": string,
      "resources": {
        "memory": string,
        "cpu": string,
      },
      "instances": {
        "count": int,
        "percentage": float,
      },
      "build":
      "run":
    ],
  },
}
```

## Collect outputs

`POST /outputs`

```javascript
{
  "runner": string,
  "run_id": string,
}
```

## Terminate Jobs

`POST /terminate`

```javascript
{
  "runner": string,
}
```

## Healthcheck

`POST /healthcheck`

```javascript
{
  "runner": string,
  "fix": bool,
}
```
