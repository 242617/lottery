# lottery

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