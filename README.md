# lottery

Lottery service instantly generates random private keys for Ethereum network, checks it's balanse and if it's not zero makes a transaction to given wallet.

It's like lottery – choose random number, check whether it's won and grab your prize!

P.S. Obviously there is no possibility to generate the existing private key or catch the collision, so it's just a joke.

## Configuration

|    Parameter     |                           Description                          |
|:----------------:|:--------------------------------------------------------------:|
| `log_prefix`     | where the log files will appear                                |
| `node_address`   | ethereum node address (or something like `infura.io`) |
| `node_secret`    | secret for basic authentication                                |
| `streams`        | number of processes in parallel                                |
| `target_address` | where to transfer funds in case of success                     |
| `gas_limit`      | gas limit for transaction                                      |

## Setup service

```
mkdir /var/log/lottery/

mkdir /opt/lottery/
cp lottery.env /opt/lottery/

cp config.yaml /etc/lottery.yaml
cp lottery /usr/local/
chmod +x /usr/local/lottery

cp lottery.service /etc/systemd/system/
systemctl enable lottery
systemctl start lottery
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