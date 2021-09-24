# Docker Settings



Since the update to Linux Kernel 5.12.2, we can no longer modify the `net.netfilter.nf_conntrack_max` sys config programmatically via the Docker API. If you encounter any network related errors, please increase it to at least 120000.

### Linux

The value `vm_map_max_count` can be permanently changed in `/etc/sysctl.conf`:

```
$ grep vm.max_map_count /etc/sysctl.conf
vm.max_map_count=120000
```

You can also change it via:

```text
sysctl -w vm.max_map_count=120000
```

### macOS via Docker for Mac

The value `vm_max_map_count` should be changed within the xhyve virtual machine:

```text
$ screen ~/Library/Containers/com.docker.docker/Data/com.docker.driver.amd64-linux/tty
```

Then, log in with _root_ and no password and configure the `sysctl` setting:

```text
sysctl -w vm.max_map_count=120000
```

### macOS via Docker Toolbox

The value `vm_max_map_count` can be changed  via docker-machine:

```text
docker-machine ssh
sudo sysctl -w vm.max_map_count=120000
```



