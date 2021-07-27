## sifgen node create

Create a new node.

```
sifgen node create [chain-id] [moniker] [mnemonic] [flags]
```

### Options

```
      --admin-clp-addresses string           admin clp addresses
      --admin-oracle-address string          admin oracle addresses
      --bind-ip-address string               IPv4 address to bind the node to (default "127.0.0.1")
      --bond-amount string                   bond amount (default "1000000000000000000000000rowan")
      --clp-config-url string                URL of the JSON file to use to pre-populate CLPs during genesis
      --enable-api                           enable API
      --enable-grpc                          enable gRPC
      --genesis-url string                   genesis URL
      --gov-max-deposit-period duration      governance max deposit period (default 15m0s)
      --gov-voting-period duration           governance voting period (default 15m0s)
  -h, --help                                 help for create
      --min-clp-create-pool-threshold uint   minimum CLP create pool threshold (default 100)
      --mint-amount string                   mint amount (default "999000000000000000000000000rowan")
      --peer-address string                  peer node to connect to
      --print-details                        print the node details
      --standalone                           standalone node
      --with-cosmovisor                      setup cosmovisor
```

### SEE ALSO

* [sifgen node](sifgen_node.md)	 - Node commands.

###### Auto generated by spf13/cobra on 2-Jul-2021