# Managing test plans

The unit for testing in the testground is a `test plan` also referred to as simply a `plan`. This document explains the layout of test plans directory in the testground home, as well is tooling for managing local plans.

### Anatomy of a test plan

A test plan is a directory which contains the code to be executed during the test and a toml-formatted manifest flile which explains the plan to the testground system. Writing new test plans from scratch is covered elsewhere in this tutorial, for now, just know that a `test plan` is a directory with a test plan manifest and that each test plan may have one or more \`test casestestground provides some basic functionalit to assist importing and creating test plans.

During execution, an archive of the plan is sent to the testground daemon, where it is built and executed. Any code or files inside the the plan directory will be available when the plan is built. 

### Where are testground testplans?

Test plans reside inside the testground home in a subdirectory called `plans`.  The location of the the testground home is governed by an envrionment variable `$TESTGROUND_HOME` and has some sane defaults if this variable is unset.

By default, on unix-like operating systems, this directory is in the user's root directory. Don't worry about creating the testground home  in advance; it will be created for you when testground runs. 

```bash
[~]↬ tree $TESTGROUND_HOME
testground
├── data
│   ├── outputs
│   └── work
├── plans           # <-- This is where plans will go!
└── sdks
```

### Plan management tooling

Test plans can be managed using regular filesystem utilities. However, the testground tool does have utilities for managing plans which can help to import and manage plans in a predictable way.  For the following sections, I'll demonstrate a few management commands along with some standard utilities to explain what the command does.

#### Creating new plans

```bash
# Create the plan
[~]↬ testground plan create -p myplan

# What happened?
[~]↬ tree $TESTGROUND_HOME
testground
├── data
│   ├── outputs
│   └── work
├── plans
│   └── myplan             # <-- the root of the plan you just created
│       ├── go.mod
│       ├── main.go        # <-- source code for your your new plan
│       └── manifest.toml  # <-- manifest for your new plan
└── sdks
```

{% hint style="info" %}
You can modify the behavior of `plan create` using commandline flags to change the module or add a git remote.
{% endhint %}

#### Importing existing plans...

Importing existing plans requires a `--source` flag. The source can be either from the local filesystem or downloaded from a git repository. When importing plans from a local filesystem, a symbolic link is created from the source to the plan directory. When git is used,  the plan is imported over any protocol supported by git.

#### ...from the local filesystem

```bash
# import multiple plans from your local filesystem
# Changing the name to "myplans" (not required)

[~]↬ testground plan import --source /local/plans/dir/ --name myplans
created symlink $TESTGROUND_HOME/plans/myplans -> /local/plans/dir/
imported plans:
myplans/verify     uses-data-network
myplans/network    ping-pong
myplans/benchmarks all
myplans/benchmarks startup
myplans/benchmarks netinit
myplans/benchmarks netlinkshape
myplans/benchmarks barrier
myplans/benchmarks subtree
myplans/placebo    ok
myplans/placebo    abort
myplans/placebo    metrics
myplans/placebo    panic
myplans/placebo    stall
myplans/example    output
myplans/example    failure
myplans/example    panic
myplans/example    params
myplans/example    sync
myplans/example    metrics

# What happened?
# a symbolic link has been created to point to the source on my local filesystem

[~]↬ ls -l $TESTGROUND_HOME/plans
total 3
drwxr-xr-x 1 cory users  60 May  4 15:01 myplan
lrwxrwxrwx 1 cory users  56 May  4 15:36 myplans -> /local/plans/dir/
```

#### ...from a git repository

```bash
# Import multiple plans from the same git repo

[~]↬ testground plan import --git --source https://github.com/libp2p/test-plans
Enumerating objects: 54, done
Counting objects: 100% (54/54), done.
Compressing objects: 100% (41/41), done.
Total 54 (delta 16), reused 36 (delta 11), pack-reused 0
cloned plan $TESTGROUND_HOME/plans/test-plans -> ssh://git@github.com/libp2p/test-plans
imported plans:
test-plans/dht find-peers
test-plans/dht find-providers
test-plans/dht provide-stress
test-plans/dht store-get-value
test-plans/dht get-closest-peers
test-plans/dht bootstrap-network
test-plans/dht all


# What happened?
# The repository was cloned with the git remote set to the source.

[~]↬ cd $TESTGROUND_HOME/plans/test-plans
[~/.../test-plans] (master=)↬ git remote -v
origin	git@github.com:libp2p/test-plans (fetch)
origin	git@github.com:libp2p/test-plans (push)
```

#### Listing plans we have added so far.

As you can see from the commands above, we have the ability to create new plans which we will write ourselves or import existing plans or collections of plans. Lets show them all with the list command.

```bash
# Generate a list of all test plans, along with all test cases in each plan.
# These are all the plans imported or created

[~]↬ testground plan list --testcases
myplan             quickstart
myplans/benchmarks all
myplans/benchmarks startup
myplans/benchmarks netinit
myplans/benchmarks netlinkshape
myplans/benchmarks barrier
myplans/benchmarks subtree
test-plans/dht     find-peers
test-plans/dht     find-providers
test-plans/dht     provide-stress
test-plans/dht     store-get-value
test-plans/dht     get-closest-peers
test-plans/dht     bootstrap-network
test-plans/dht     all
myplans/example    output
myplans/example    failure
myplans/example    panic
myplans/example    params
myplans/example    sync
myplans/example    metrics
myplans/network    ping-pong
myplans/placebo    ok
myplans/placebo    abort
myplans/placebo    metrics
myplans/placebo    panic
myplans/placebo    stall
myplans/verify     uses-data-network
```

#### Removing plans

Finally, lets end by removing a plan we are no longer interested in.

```bash
# Examples? Who needs em!

[~]↬ testground plan rm -p myplans/example

# Oops! This command is destructive, so it needs a confirmation.

[~]↬ testground plan rm -p myplans/example --yes
```



