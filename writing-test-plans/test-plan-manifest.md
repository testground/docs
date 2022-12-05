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

The next few sections are options passed to each builder and runner when they are selected. Make sure to enable any [builders](../concepts-and-architecture/builders.md#supported-builders) and [runners](../concepts-and-architecture/runners.md#supported-runners) you want to use! The following section enables all runners and passes configurations. Of particular interest is the module path. Make sure this is correct to ensure code can be correctly imported.

```toml
[builders."docker:go"]
enabled = true

[builders."exec:go"]
enabled = true

[runners."local:docker"]
enabled = true

[runners."local:exec"]
enabled = true

[runners."cluster:k8s"]
enabled = true
```

### Test cases

Finally, we have [test cases](../concepts-and-architecture/test-structure.md#test-cases). Test cases are defined in an [array of tables](https://github.com/toml-lang/toml#array-of-tables) which specify the name of the test case, boundaries for the number of instances and the values of any parameters being tested in a particular test case.
If your plan requires an integer param without a default value, the test will fail if you do not provide it.

```toml
[[testcases]]
name= "quickstart"
instances = { min = 1, max = 5, default = 1 }

[testcases.params]
  conn_count       = { type = "int", desc = "number of TCP sockets to open" default = 5 }
  conn_outgoing    = { type = "int", desc = "number of outgoing TCP dials", default = 5 }
  conn_delay_ms       = { type = "int", desc = "random milliseconds jitter before TCP dial", default = 30000 }
  concurrent_dials = { type = "int", desc = "max number of concurrent net.Dial calls", default = 10 }
  data_size_kb     = { type = "int", desc = "size of data to write to each TCP connection", default = 128 }
  barrier_iterations = { type = "int", desc = "number of iterations of the barrier test", unit = "iteration", default = 10 }
  barrier_test_timeout_secs  = { type = "int", desc = "barrier testcase timeout", unit = "seconds", default = 300 }
  subtree_iterations = { type = "int", desc = "number of iterations of the subtree test", unit = "iteration", default = 2000 }
  subtree_test_timeout_secs  = { type = "int", desc = "subtree testcase timeout", unit = "seconds", default = 300 }
  expected_version = { type = "string", desc = "expected version" }
  expected_implementation = { type = "string", desc = "expected implementation" }
  messages   = { type = "int", default = 50 }
  ...

```

### The resulting `manifest.toml`  test plan manifest

Putting it all together:

```toml
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
