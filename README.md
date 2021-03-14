# lottery

Lottery service instantly generates random private keys for Ethereum network, checks it's balanse and if it's not zero makes a transaction to given wallet.

It's like a lottery – choose random number, check whether it's won and grab your prize!

P.S. Obviously there is no possibility to generate the existing private key or catch the collision, so it's just a joke.

## Configuration

```
ethereum:
  node_address: http://localhost:8070
  gas_limit: 21000
streams: 10
target_address: 0xa1a831aa268797016af2061c5bb72775bcaf40ee
stat_period: 10s
```

## Utils

# Private key to address

Returns your address for given private key.

```
# build
make utils:private2address

# run
./private2address {{private_key}}
```

# Extract your private key from keystore file

Returns your private key for given keystore file (e.g. UTC--2019-10-23T14-0...).

```
# build
make utils:keystore2private

# run
./keystore2private {{keystore_file}}
```
