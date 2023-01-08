#!/bin/bash

curl -LO https://github.com/prometheus/node_exporter/releases/download/v1.2.2/node_exporter-1.2.2.linux-386.tar.gz
tar -xvzf $HOME/node_exporter-1.2.2.linux-386.tar.gz
sudo cp $HOME/node_exporter-1.2.2.linux-386/node_exporter /usr/local/bin/
sudo useradd node_exporter
sudo chown node_exporter:node_exporter /usr/local/bin/node_exporter
rm node_exporter-1.2.2.linux-386* -rf

echo -e "[Unit]\nDescription=Node Exporter\nWants=network-online.target\nAfter=network-online.target\n\n[Service]\nUser=node_exporter\nGroup=node_exporter\nType=simple\nExecStart=/usr/local/bin/node_exporter\n\n[Install]\nWantedBy=multi-user.target" | sudo tee -a /etc/systemd/system/node_exporter.service

sudo systemctl daemon-reload
sudo service node_exporter start
bash <(curl -Ss https://my-netdata.io/kickstart.sh)
