# Picking a runner

Plans which run on one runner generally should run be runnable on all other runners as well. The following table describes the features of different runners.

It is common practice when developing a test plan to use a local runner \(`local:exec` or `local:docker`\) in order to iterate quickly and then move to the Kubernetes `cluster:k8s` runner when you want to run your test plan with many more test instances.

|  runner | quick iteration | high instance count | network containment | quick setup |
| :--- | :--- | :--- | :--- | :--- |
| **local:exec** | ✅ | ❌ | ❌ | ✅ |
| **local:docker** | ✅ | ❌ | ✅ | ✅ |
| **cluster:k8s** | ❌ | ✅ | ✅ | ❌ |

