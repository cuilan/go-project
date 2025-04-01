#!/bin/bash

set -e

cp ./your-go-project.service /usr/lib/systemd/system/
cp ./your-go-project.service.env /usr/lib/systemd/system/
cp your-go-project /usr/local/bin/

systemctl daemon-reload
systemctl start your-go-project.service
systemctl enable your-go-project.service
#systemctl restart your-go-project.service

systemctl status your-go-project.service
