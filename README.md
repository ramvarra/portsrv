# HTTPServers on multiple ports

`portsrv` creates HTTP servers on a list of ports.
Can be used for testing if the ports are accessible on host from other hosts on the
network.

## Build
```
    $ go build 
```
## Usage:

To start HTTP Servers on ports8181, 8182, 8183 on host srv

```
srv $> portsrv 8181 8182 8183

```

To test from a client

```
client $> for p in 8181 8182 8183; do curl srv:$p; done
```
