# exec:go builder

The `exec:go` builder uses the user's own go installation to compile and build a binary. Using this builder will use and alter the user's go pkg cache. None of these are required and need only be edited if the defaults do not work well in your environment.

[github](https://github.com/ipfs/testground/blob/master/pkg/build/golang/exec.go#L28)

| parameter | explanation |
| :--- | :--- |
| module\_path | use an alternative gomod path |
| exec\_pkg | specify the package name |
| fresh\_gomod | remove and recreate `go.mod` files |

