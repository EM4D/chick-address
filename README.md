# Chick-address

Chick-address is a service for viewing the exact IP address along with several other IP-related information 


The **[ip-api.com]("https://ip-api.com")** website API is currently in use

## Usage
```
./chick-address -port=8080 -url="yourdomain.com"
```

## Assets
To use the assets first, create a directory (if not exists) at the root of the project called assets, then put the files you need in it.

**Your assets are available at example.com/assets/**

## Run as daemon
Create a daemon file for example: /etc/systemd/system/chick-address.service with below sample
```
[Unit]
Description=chick-address
Wants=network-online.target
After=network-online.target

[Service]
User=ubuntu
Group=ubuntu
Type=simple
WorkingDirectory=/home/ubuntu/chick-address/
ExecStart=/home/ubuntu/chick-address/chick-address

[Install]
WantedBy=multi-user.target
```
Then:
$sudo systemctl daemon-reload
$sudo systemctl start chick-address
