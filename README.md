# Imperator
**Imperator** is a blockchain built using Cosmos SDK and Tendermint

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
imperatord init $MONIKER
```

 
 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.imperator/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.imperator/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
imperatord add-consumer-section
```

## Launch node

```
imperatord start
```
