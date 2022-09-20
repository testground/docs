# Understanding the test plan manifest

Inside the root of every test plan is a file called `manifest.toml`.  This file is used to explain the test plan to the Testground daemon. In this file,  a plan can be restricted to specific runners. Additionally, this is the file where test cases and parameters are defined.

If you have created the `quickstart` plan in the previous article, let's have a look at the manifest file we have just created.

### Test plan name

the name of the test plan is a top-level of the file's schema. 

```text
name = "quickstart"
```

### Defaults

Enclosed in a defaults section are default builders. These are variables which can be used as defaults for the plan In this section, we set the default builder and runner.

```text
[defaults]
builder = "exec:go"
runner = "local:exec"
```

### Builder and runner options

The next few sections are options passed to each builder and runner when they are selected. Make sure to enable any [builders](../concepts-and-architecture/builders-1.md#supported-builders) and [runners](../concepts-and-architecture/runners.md#supported-runners) you want to use! The following section enables all runners and passes configurations. Of particular interest is the module path. Make sure this is correct to ensure code can be correctly imported.

```yaml
[builders."docker:go"]
enabled = true
go_version = "1.14"
module_path = "github.com/your/module/name"
exec_pkg = "."
fresh_gomod = true
modfile = "./go.mod" # Custom modfile

#  GoProxyMode specifies one of "local", "direct", "remote".
# 
#    * The "local" mode (default) will start a proxy container (if one
#      doesn't exist yet) with bridge networking, and will configure the
#      build to use that proxy.
#    * The "direct" mode sets the `GOPROXY=direct` env var on the go build.
#    * The "remote" mode specifies a custom proxy. The `GoProxyURL` field
#      must be non-empty.
go_proxy_mode = "local"

# GoProxyURL specifies the URL of the proxy when GoProxyMode = "custom".
go_proxy_url = ""

# RuntimeImage is the runtime image that the test plan binary will be
# copied into. Defaults to busybox:1.31.1-glibc.
runtime_image = "busybox:1.31.1-glibc"

# BuildBaseImage is the base build image that the test plan binary will be
# built from. Defaults to golang:1.16-buster
build_base_image = "golang:1.16-buster"

# SkipRuntimeImage allows you to skip putting the build output in a
# slimmed-down runtime image. The build image will be emitted instead.
skip_runtime_image = true

# EnableGoBuildCache enables the creation of a go build cache and its usage.
# When enabling for the first time, a cache image will be created with the
# dependencies of the current plan state.
# If this flag is unset or false, every build of a test plan will start
# with a blank go container. If this flag is true, the builder will the last
# cached image.
enable_go_build_cache = false

#  Cgo enables the creation of Go packages that call C code. By default it is disabled.
#  Enabling CGO also enables dynamic linking. Disabling CGO (default) produces statically  linked binaries.
#  If you ever see errors like the following, you are probably better off
#  disabling CGO (and therefore enabling static linking).
#    /testplan: error while loading shared libraries: libdl.so.2: cannot open shared object file: No such file or directory
#  If you pass `true` to this flag, your test plan will be built with CGO_ENABLED=1
enable_cgo = false

[builders."exec:go"]
enabled = true
module_path = "github.com/your/module/name"

[runners."local:docker"]
enabled = true
ulimits = [
  "nofile=1048576:1048576",
]

[daemon]
listen                    = ":8040"

[daemon.scheduler]
task_timeout_min          = 5
task_repo_type            = "disk"

[client]
endpoint = "http://localhost:8040"
user = "myname"

[runners."local:exec"]
enabled = true

[runners."cluster:k8s"]
enabled = true
run_timeout_min             = 1

# Resources requested for each testplan pod from the Kubernetes cluster
testplan_pod_cpu            = "10m"
testplan_pod_memory         = "10Mi"

#  Resources requested for the `collect-outputs` pod from the Kubernetes cluster
collect_outputs_pod_cpu     = "10m"
collect_outputs_pod_memory  = "10Mi"
autoscaler_enabled          = false
provider                    = ""
sysctls = []

```

### Test cases

Finally, we have [test cases](../concepts-and-architecture/test-structure.md#test-cases). Test cases are defined in an [array of tables](https://github.com/toml-lang/toml#array-of-tables) which specify the name of the test case, boundaries for the number of instances and the values of any parameters being tested in a particular test case.

```yaml
[[testcases]]
name= "quickstart"
instances = { min = 1, max = 5, default = 1 }

[testcases.params]
  conn_count       = { type = "int", desc = "number of TCP sockets to open" default = 5 }
  conn_outgoing    = { type = "int", desc = "number of outgoing TCP dials", default = 5 }
  conn_delay_ms       = { type = "int", desc = "random milliseconds jitter before TCP dial", default = 30000 }
  concurrent_dials = { type = "int", desc = "max number of concurrent net.Dial calls", default = 10 }
  data_size_kb     = { type = "int", desc = "size of data to write to each TCP connection", default = 128 }
```

### The resulting `manifest.toml`  test plan manifest

Putting it all together:

```yaml
name = "quickstart"

[defaults]
builder = "exec:go"
runner = "local:exec"

[builders."docker:go"]
enabled = true
go_version = "1.14"
module_path = "github.com/your/module/name"
exec_pkg = "."

[builders."exec:go"]
enabled = true
module_path = "github.com/your/module/name"

[runners."local:docker"]
enabled = true

[runners."local:exec"]
enabled = true

[runners."cluster:k8s"]
enabled = true

[[testcases]]
name= "quickstart"
instances = { min = 1, max = 5, default = 1 }

# Add more testcases here...
# [[testcases]]
# name = "another"
# instances = { min = 1, max = 1, default = 1 }
#   [testcases.params]
#   param1 = { type = "int", desc = "an integer", unit = "units", default = 3 }

```



