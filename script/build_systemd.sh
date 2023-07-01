#!/bin/bash

systemctl stop remote-part-job

# check if the remote-part-job service is running
active=`systemctl is-active remote-part-job`
if [ $active = "active" ]; then
    echo "build failed"
    echo "remote-part-job service is still active, please use 'systemctl stop remote-part-job.service' to stop the service"
    exit -1
fi

# get the script abs filepath
folder=$(dirname $(readlink -f "$0"))
preFolder=$(dirname "$folder")

cat << EOF > "remote-part-job.service"
[Unit]
Description=remote-part-job
After=network.target

[Service]
Type=forking
PIDFile=$preFolder/pidfile.txt
Restart=always
RestartSec=10
StartLimitInterval=0
LimitNOFILE=102400
LimitNPROC=102400
ExecStart=/bin/bash $folder/start.sh start

[Install]
WantedBy=multi-user.target
EOF

cp -f remote-part-job.service /etc/systemd/system/
rm -rf remote-part-job.service
systemctl daemon-reload
systemctl enable remote-part-job.service
systemctl start remote-part-job.service

# check if the remote-part-job service is running
active=`systemctl is-active remote-part-job`
if [ $active != "active" ]; then
  echo "start failed"
  exit -1
fi

echo "start remote-part-job success"