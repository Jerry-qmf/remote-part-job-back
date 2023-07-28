#!/bin/bash

set -e

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
ROOT="$( dirname "$DIR" )"
cd $ROOT

PROGRAM=bin/remote-part-job
LOGS=/data/log/remote-part-job

function Usage() {
  echo "Usage: $0 [start|stop|restart]"
  echo "example: $0 start"
  exit 0
}

function Start() {

  if [ ! -d "$LOGS" ]; then
    mkdir -p $LOGS
  fi

  #RUN
  export GODEBUG=madvdontneed=1
  nohup $PROGRAM > $LOGS/error.log 2>&1 &
  echo $! > pidfile.txt
  echo "start success!"
}

function Restart() {
  Stop
  Start
}

function Check() {
  pid=$(ps -ef | grep $PROGRAM | grep -v grep | awk '{print $2}')
  if [ -n "$pid" ]; then
    exit 0
  fi

  echo "restart $PROGRAM"
  Start
}

function Stop() {
  pid=$(ps -ef | grep $PROGRAM | grep -v grep | awk '{print $2}')
  if [ -n "$pid" ]; then
    kill -9 $pid
    rm -f pidfile.txt
    echo "stop success pid=$pid"
  else
    echo "$PROGRAM has stopped!"
  fi
}

case $1 in
  start) Start $@;;
  restart) Restart $@;;
  check) Check $@;;
  stop) Stop $@;;
  *) Usage $@;;
esac