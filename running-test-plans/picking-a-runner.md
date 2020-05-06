# Picking a runner

Plans which run on one runner will run on them all. The following describes the features of different runners. 

It is a common practice when developing a plan to use a local runner to iterate quickly and then move to the kubernetes runner when the time comes to run a larger test.



|  runner | quick iteration | high instance count | network containment | quick setup |
| :--- | :--- | :--- | :--- | :--- |
| **local:exec** | ✅ | ❌ | ❌ | ✅ |
| **local:docker** | ✅ | ❌ | ✅ | ✅ |
| **cluster:k8s** | ❌ | ✅ | ✅ | ❌ |



#### future development:

* Integration with CI/CD
* Automated analysis pipelines

