# dolan

[![Join the chat at https://gitter.im/tam7t/dolan](https://badges.gitter.im/tam7t/dolan.svg)](https://gitter.im/tam7t/dolan?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## About

This utility helps secure the `private` interface on DigitalOcean droplets by
adding `iptable` rules that only allow traffic from your other droplets. `dolan`
queries the DigitalOcean API and automatically updates `iptable` rules.

## Usage

```
DO_KEY=<read_only_api_token> /path/to/dolan_linux_amd64
```

The `iptables` rules added by `dolan` are equivalent to:

```
-N dolan-peers # create a new chain
-A INPUT -i eth1 -j dolan-peers # add chain to private interface
-A INPUT -i eth1 -j DROP # add default DROP rule to private interface
-A dolan-peers -s <PEER>/32 -j ACCEPT # allow traffic from PEER ip address
```

## Development

### Dependencies

Dependencies are vendored with [govendor](https://github.com/kardianos/govendor).

### Build

A `Makefile` is included:
  * `test` - runs unit tests
  * `build` - builds `dolan` for all platforms
  * `build_local` - builds `dolan` for just the local architecture
