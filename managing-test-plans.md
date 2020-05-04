# Managing test plans

The unit for testing in the testground is a `test plan` also referred to as simply a `plan`. This document explains the layout of test plans as well is tooling for managing local plans.

### Where are testplans?

Test plans reside inside the testground home in a subdirectory called `plans`.  The location of the the testground home is governed by an envrionment variable `$TESTGROUND_HOME` and has some sane defaults if this variable is unset.

By default, on unix-like operating systems, this directory is in the user's root directory. Don't worry about creating the testground home  in advance; it will be created for you when testground runs. 

```text
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

```text
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

#### Importing existing plans

Importing existing plans requires a `--source` flag. The source can be either from the local filesystem or downloaded from a git repository. When importing plans from a local filesystem, a symbolic link is created from the source to the plan directory. When git is used,  the plan is imported over any protocol supported by git.

```text
# Import multiple plans from the same git repo

[~]↬ testground plan import --git --source https://github.com/libp2p/test-plans
Enumerating objects: 54, done
Counting objects: 100% (54/54), done.
Compressing objects: 100% (41/41), done.
Total 54 (delta 16), reused 36 (delta 11), pack-reused 0
cloned plan /home/cory/testground/plans/test-plans -> ssh://git@github.com/libp2p/test-plans
imported plans:
test-plans/dht find-peers
test-plans/dht find-providers
test-plans/dht provide-stress
test-plans/dht store-get-value
test-plans/dht get-closest-peers
test-plans/dht bootstrap-network
test-plans/dht all


# What happened? The repository was cloned with the git remote set to the source.

[~]↬ cd $TESTGROUND_HOME/plans/test-plans
[~/.../test-plans] (master=)↬ git remote -v
origin	git@github.com:libp2p/test-plans (fetch)
origin	git@github.com:libp2p/test-plans (push)
```



* * Where test plans must be located locally \($TESTGROUND\_HOME/plans\).
* Importing a test plan from a Git repo.
* Importing a test plan from a local folder
* Listing enrolled test plans
* Removing a test plan



