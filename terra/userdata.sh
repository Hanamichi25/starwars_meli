#!/bin/bash
export PATH=$PATH:/usr/local/bin
yum install unzip;
aws s3api get-object --bucket starwars-mercadolibre --key stawars.zip /tmp/stawars.zip;
cd "/tmp";
unzip -o stawars.zip;
echo "[Unit]
Description = making network connection up
After = network.target

[Service]
ExecStart = /tmp/starwars

[Install]
WantedBy = multi-user.target" > /tmp/starwars.service
cp /tmp/starwars.service /usr/lib/systemd/system/
systemctl enable starwars.service
systemctl start starwars.service