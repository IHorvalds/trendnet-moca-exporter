# trendnet-moca-exporter
(Simplistic) Prometheus exporter for TrendNet MoCA 2.5 Adapters

This *only* monitors the state of the MoCA connection. Obviously, if the MoCA connection is the only connection for one of the devices to the IP network, that adapter won't respond until the MoCA connection comes back up 😀. That's fine, you have another one which must be connected to some kind of IP network.

# Configuration file format
By default, the docker image will load `/configs/moca.toml`

Its expected format is:
```toml
["name-of-adapter-1"]
AdapterAddress = "ip.of.moca.adapter"
User = "admin-username"
Password = "admin-password"

["name-of-adapter-2"]
...
```
