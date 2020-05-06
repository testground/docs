# Understanding the test plan manifest

Inside the root of each test plan is a file called `manifest.toml`.  This file is used to explain the test plan to the testground daemon. In this file,  a plan can be restricted to specific runners. Additionally, this is the file where testcases and parameters are defined.

If you have created the quickstart plan in the previous article, let's have a look at the manifest file we have just created.  

#### test plan name

the name of the test plan is a top-level of the file's schema. 

```text
name = "quickstart"
```

#### defaults

Enclosed in a defaults section are default builders. These are variables which can be used as defaults for the plan In this section, we set the default builder and runner.

```text
[defaults]
builder = "exec:go"
runner = "local:exec"
```

#### builder and runner options

The next few sections are options passed to each builder and runner when they are selected. Make sure to enable any runners you want to use! The following section enables all runners and passes configurations. Of particular interest is the module path. Make sure this is correct to ensure code can be correctly imported.

```text
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
```

#### testcases

Finally, we have test cases. Testcases are defined in an [array of tables](https://github.com/toml-lang/toml#array-of-tables) which specify the name of the testcase, boundaries for the number of instances and the values of any parameters being tested in a particular test case.

```text
[[testcases]]
name= "quickstart"
instances = { min = 1, max = 5, default = 1 }
```



#### putting it all together:

```text
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
#   [testcase.params]
#   param1 = { type = "int", desc = "an integer", unit = "units", default = 3 }

```



