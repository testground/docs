# How does it work?

### 1. **You develop distributed test plans as if you were writing unit tests against local APIs**

* No puppeteering necessary.
* No need to package and ship the system as a separate daemon with an external API in order to puppeteer it.
* No need to expose every internal setting over an external API, just for the sake of testing.

### **2. Your test plan calls out to the coordination API to...**

* communicate out-of-band information \(such as endpoint addresses, peer ids, etc.\)
* leverage synchronization and ordering primitives such as signals and barriers to model a distributed state machine.
* programmatically apply network traffic shaping policies, which you can alter during the execution of a test to simulate various network conditions.

### **3. There is no special "conductor" node telling instances what to do when**

* The choreography and sequencing emerges from within the test plan itself.

### **4. You decide what versions of the upstream software you want to exercise your test against.**

* Benchmark, simulate, experiment, run attacks, etc. against versions v1.1 and v1.2 of the components under test in order to compare results, or test compatibility.
* Assemble hybrid test runs mixing various versions of the dependency graph.

### **5. Inside your test plan...**

* You record observations, metrics, success/failure statuses.
* You emit structured or unstructured assets you want collected, such as event logs, dumps, snapshots, binary files, etc.

### **6. Via a TOML-based** _**composition**_ **file, you instruct Testground to...**

* Assemble a test run comprising groups of 2, 200, or 10000 instances, each with different test parameters, or built against different depencency sets.
* Schedule them for run locally \(executable or Docker\), or in a cluster \(Kubernetes\).

### **7. You collect the outputs of the test plan for analysis...**

* **with a single command.** 
* use data processing scripts and platforms \(such as the upcoming Jupyter notebooks integration\) to draw conclusions.

